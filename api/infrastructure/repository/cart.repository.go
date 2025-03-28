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

func (r *cartRepositoryImpl) RemoveItemFromCart(ctx context.Context, ownerID int64, ID int64) error {
	var cartItem domain.CartItem
	if err := r.db.Where("id = ?", ID).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrCartItemNotFound
		}
		return err
	}
	if cartItem.OwnerID != ownerID {
		return domain.ErrCartItemNotBelongToOwner
	}
	err := r.db.Delete(&cartItem).Error
	return err
}

func (r *cartRepositoryImpl) ClearItemsByOwnerID(ctx context.Context, ownerID int64) error {
	return r.db.Where("owner_id = ?", ownerID).Delete(&domain.CartItem{}).Error
}

func (r *cartRepositoryImpl) GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) ([]*domain.CartItem, error) {
	var cartItems []*domain.CartItem
	err := r.db.Preload("ArtToy.Owner").Where("owner_id = ?", ownerID).Find(&cartItems).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*domain.CartItem{}, nil
	}
	if err != nil {
		return nil, err
	}
	return cartItems, err
}

var _ domain.CartRepository = &cartRepositoryImpl{}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepositoryImpl{
		db,
	}
}
