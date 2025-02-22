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

func (r *cartRepositoryImpl) GetCartByOwnerID(ctx context.Context, ownerID int64) ([]*domain.ArtToy, error) {
	var cartItems []domain.CartItem
	var cart domain.Cart
	
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCart := domain.NewCart(ownerID)
		err := r.CreateCart(ctx, newCart)
		if err != nil {
			return nil, err
		}
	}
	err = r.db.WithContext(ctx).
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.owner_id = ?", ownerID).
		Preload("ArtToy").
		Find(&cartItems).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrCartNotFound
	}

	if len(cartItems) == 0 {
        return []*domain.ArtToy{}, err
    }

	var artToys []*domain.ArtToy
	for _, item := range cartItems {
		artToys = append(artToys, &item.ArtToy)
	}
	return artToys, err
}

var _ domain.CartRepository = &cartRepositoryImpl{}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepositoryImpl{
		db,
	}
}
