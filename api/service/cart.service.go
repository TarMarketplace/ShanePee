package service

import (
	"context"
	"errors"

	"shanepee.com/api/domain"
)

var (
	ErrCartItemNotFound error = domain.ErrCartItemNotFound
)

type CartService interface {
	AddItemToCart(ctx context.Context, ownerID int64, artToyID int64) (*domain.CartItem, error)
	RemoveItemFromCart(ctx context.Context, ownerID int64, artToyID int64) error
	GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) ([]*domain.CartItem, error)
	Checkout(ctx context.Context, ownerID int64) error
}

type cartServiceImpl struct {
	artToyRepo domain.ArtToyRepository
	cartRepo   domain.CartRepository
	orderRepo  domain.OrderRepository
}

func NewCartService(artToyRepo domain.ArtToyRepository, cartRepo domain.CartRepository, orderRepo domain.OrderRepository) CartService {
	return &cartServiceImpl{
		artToyRepo,
		cartRepo,
		orderRepo,
	}
}

var _ CartService = &cartServiceImpl{}

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, ownerID int64, artToyID int64) (*domain.CartItem, error) {
	cartItem := domain.NewCartItem(ownerID, artToyID)
	err := s.cartRepo.AddItemToCart(ctx, cartItem)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (s *cartServiceImpl) RemoveItemFromCart(ctx context.Context, ownerID int64, ID int64) error {
	err := s.cartRepo.RemoveItemFromCart(ctx, ownerID, ID)
	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
			return ErrCartItemNotFound
		}
		return err
	}
	return nil
}

func (s *cartServiceImpl) GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) ([]*domain.CartItem, error) {
	cart, err := s.cartRepo.GetCartWithItemByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *cartServiceImpl) Checkout(ctx context.Context, ownerID int64) error {
	cartItems, err := s.cartRepo.GetCartWithItemByOwnerID(ctx, ownerID)
	if err != nil {
		return err
	}

	// All items should be owned by the same seller
	sellerID := cartItems[0].ArtToy.OwnerID
	order := domain.NewOrder(sellerID, ownerID)
	if err := s.orderRepo.CreateOrder(ctx, order); err != nil {
		return err
	}

	orderItems := make([]*domain.OrderItem, 0)
	artToyIDs := make([]int64, 0)
	for _, cartItem := range cartItems {
		artToyIDs = append(artToyIDs, cartItem.ArtToyID)

		if !cartItem.ArtToy.Availability {
			return ErrArtToyNotFound
		}
		orderItem := domain.NewOrderItem(cartItem.ArtToyID, order.ID)
		orderItems = append(orderItems, orderItem)
	}

	if err := s.orderRepo.CreateOrderItems(ctx, orderItems); err != nil {
		return err
	}
	if err := s.artToyRepo.UpdateArtToysAvailability(ctx, artToyIDs, false); err != nil {
		return err
	}

	// TODO: implement DeleteItemsInCart and handle error
	// err = s.cartRepo.DeleteItemsInCart(ctx, cartID)
	// if err != nil {
	// 	return err
	// }

	// TODO: make transaction
	return nil
}
