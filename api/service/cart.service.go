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
	CreateCart(ctx context.Context, ownerID int64) (*domain.Cart, error)
	GetCartByOwnerID(ctx context.Context, ownerID int64) ([]*domain.ArtToy, error)
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

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, cartID int64, artToyID int64) (*domain.CartItem, error) {
	cartItem := domain.NewCartItem(cartID, artToyID)
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

func (s *cartServiceImpl) GetCartByOwnerID(ctx context.Context, ownerID int64) ([]*domain.ArtToy, error) {
	cart, err := s.cartRepo.GetCartByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
