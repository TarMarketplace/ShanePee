package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type cartRepositoryImpl struct {
	db *gorm.DB
}

func (r *cartRepositoryImpl) AddItemToCart(ctx context.Context, cartItem *domain.CartItem) error {
	err := r.db.Create(cartItem).Error

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return domain.ErrArtToyNotFound
	}
	return err
}

func (r *cartRepositoryImpl) GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.Preload("Items").Where("owner_id = ?", ownerID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCart := domain.NewCart(ownerID)
		err := r.CreateCart(ctx, newCart)
		if err != nil {
			return nil, err
		}
		return newCart, nil
	}
	return &cart, err
}

var _ domain.CartRepository = &cartRepositoryImpl{}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepositoryImpl{
		db,
	}
}
