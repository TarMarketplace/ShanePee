package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrCartNotFound          error = domain.ErrCartNotFound
	ErrCartAndArtToyNotFound error = domain.ErrCartAndArtToyNotFound
)

type CartService interface {
	AddItemToCart(ctx context.Context, cartID int64, artToyID int64) (*domain.CartItem, error)
	CreateCart(ctx context.Context, sellerID int64) (*domain.Cart, error)
	Checkout(ctx context.Context, cartID int64, buyerID int64) error
}

type cartServiceImpl struct {
	cartRepo  domain.CartRepository
	orderRepo domain.OrderRepository
}

func NewCartService(cartRepo domain.CartRepository, orderRepo domain.OrderRepository) CartService {
	return &cartServiceImpl{
		cartRepo,
		orderRepo,
	}
}

var _ CartService = &cartServiceImpl{}

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, cartID int64, artToyID int64) (*domain.CartItem, error) {
	cartItem := domain.NewCartItem(cartID, artToyID)
	err := s.cartRepo.AddItemToCart(ctx, cartItem)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (s *cartServiceImpl) CreateCart(ctx context.Context, sellerID int64) (*domain.Cart, error) {
	cart := domain.NewCart(sellerID)
	err := s.cartRepo.CreateCart(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *cartServiceImpl) Checkout(ctx context.Context, cartID int64, buyerID int64) error {
	var cart *domain.Cart
	// TODO: implement GetCartByID and handle error
	// cart, err := s.cartRepo.GetCartByID(ctx, cartID)
	// if err != nil {
	// 	return err
	// }

	var items []*domain.CartItem
	// TODO: implement GetItemsInCart and handle error
	// items, err := s.cartRepo.GetItemsInCart(ctx, cartID)
	// if err != nil {
	// 	return err
	// }

	order := domain.NewOrder(cart.SellerID, buyerID)
	if err := s.orderRepo.CreateOrder(ctx, order); err != nil {
		return err
	}
	for _, item := range items {
		orderItem := domain.NewOrderItem(item.ArtToyID, order.ID)
		if err := s.orderRepo.CreateOrderItem(ctx, orderItem); err != nil {
			return err
		}
	}

	// TODO: implement DeleteItemsInCart and handle error
	// err = s.cartRepo.DeleteItemsInCart(ctx, cartID)
	// if err != nil {
	// 	return err
	// }

	// TODO: make transaction
	return nil
}
