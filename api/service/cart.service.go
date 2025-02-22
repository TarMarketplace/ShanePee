package service

import (
	"context"

	"shanepee.com/api/domain"
)

type CartService interface {
	AddItemToCart(ctx context.Context, artToyID int64, ownerID int64) (*domain.CartItem, error)
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

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, artToyID int64, ownerID int64) (*domain.CartItem, error) {
	cartItem := domain.NewCartItem(artToyID, ownerID)
	err := s.cartRepo.AddItemToCart(ctx, cartItem)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (s *cartServiceImpl) Checkout(ctx context.Context, ownerID int64) error {
	var items []*domain.CartItem
	// TODO: implement GetItemsInCart and handle error
	// items, err := s.cartRepo.GetItemsInCart(ctx, ownerID)
	// if err != nil {
	// 	return err
	// }

	// All items should be owned by the same seller
	sellerID := items[0].OwnerID
	order := domain.NewOrder(sellerID, ownerID)
	if err := s.orderRepo.CreateOrder(ctx, order); err != nil {
		return err
	}
	for _, item := range items {
		orderItem := domain.NewOrderItem(item.ArtToyID, order.ID)
		if err := s.orderRepo.CreateOrderItem(ctx, orderItem); err != nil {
			return err
		}
		if err := s.artToyRepo.UpdateArtToy(ctx, item.ArtToyID, map[string]any{
			"availability": true,
		}); err != nil {
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
