package service

import (
	"context"

	"shanepee.com/api/domain"
)

type CartService interface {
	AddItemToCart(ctx context.Context, userID int64, artToyID int64) (*domain.CartItem, error)
	CreateCart(ctx context.Context, ownerID int64) (*domain.Cart, error)
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

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, userID int64, artToyID int64) (*domain.CartItem, error) {
	cartItem := domain.NewCartItem(userID, artToyID)
	err := s.cartRepo.AddItemToCart(ctx, cartItem)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (s *cartServiceImpl) CreateCart(ctx context.Context, ownerID int64) (*domain.Cart, error) {
	cart := domain.NewCart(ownerID)
	err := s.cartRepo.CreateCart(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
