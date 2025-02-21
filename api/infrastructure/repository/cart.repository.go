package repository

import (
	"context"
	"strings"

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
	// gorm.ErrForeignKeyConstraintFailed does not exist in this gorm version
	if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
		cartNotFoundErr := r.db.First(&domain.Cart{}, cartItem.CartID).Error
		artToyNotFoundErr := r.db.First(&domain.ArtToy{}, cartItem.ArtToyID).Error
		if cartNotFoundErr != nil && artToyNotFoundErr != nil {
			return domain.ErrCartAndArtToyNotFound
		} else if err := r.db.First(&domain.Cart{}, cartItem.CartID).Error; err != nil {
			return domain.ErrCartNotFound
		} else if err := r.db.First(&domain.ArtToy{}, cartItem.ArtToyID).Error; err != nil {
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
