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

func (r *cartRepositoryImpl) CreateCart(ctx context.Context, cart *domain.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepositoryImpl) AddItemToCart(ctx context.Context, cartItem *domain.CartItem) error {
	err := r.db.Create(cartItem).Error

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		cartNotFoundErr := r.db.First(&domain.Cart{}, cartItem.CartID).Error
		artToyNotFoundErr := r.db.First(&domain.ArtToy{}, cartItem.ArtToyID).Error
		if errors.Is(cartNotFoundErr, gorm.ErrRecordNotFound) && errors.Is(artToyNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrCartAndArtToyNotFound
		}
		if errors.Is(cartNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrCartNotFound
		}
		if errors.Is(artToyNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrArtToyNotFound
		}
	}
	return err
}

var _ domain.CartRepository = &cartRepositoryImpl{}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepositoryImpl{
		db,
	}
}
