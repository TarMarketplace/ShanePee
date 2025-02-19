package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type cartRepositoryImpl struct {
	db *gorm.DB
}

func (r *cartRepositoryImpl) CreateCart(ctx context.Context, cart *domain.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepositoryImpl) AddItemToCart(ctx context.Context, cartItem *domain.CartItem) error {
	return r.db.Create(cartItem).Error
}

var _ domain.CartRepository = &cartRepositoryImpl{}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepositoryImpl{
		db,
	}
}
