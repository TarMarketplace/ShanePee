package service

import (
	"context"

	"shanepee.com/api/domain"
)

type CartService interface {
	AddItemToCart(ctx context.Context, ownerID int64, artToyID int64) (*domain.CartItem, error)
}

type cartServiceImpl struct {
	cartRepo domain.CartRepository
}

func NewCartService(cartRepo domain.CartRepository) CartService {
	return &cartServiceImpl{
		cartRepo,
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
