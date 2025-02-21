package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func (r *orderRepositoryImpl) CreateOrder(ctx context.Context, order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepositoryImpl) CreateOrderItem(ctx context.Context, orderItem *domain.OrderItem) error {
	err := r.db.Create(orderItem).Error

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		orderNotFoundErr := r.db.First(&domain.Order{}, orderItem.OrderID).Error
		artToyNotFoundErr := r.db.First(&domain.ArtToy{}, orderItem.ArtToyID).Error
		if errors.Is(orderNotFoundErr, gorm.ErrRecordNotFound) && errors.Is(artToyNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrOrderAndArtToyNotFound
		}
		if errors.Is(orderNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrOrderNotFound
		}
		if errors.Is(artToyNotFoundErr, gorm.ErrRecordNotFound) {
			return domain.ErrArtToyNotFound
		}
	}
	return err
}

func (r *orderRepositoryImpl) FindOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*domain.Order, error) {
	var order []*domain.Order
	if err := r.db.Where("seller_id = ? AND status = ?", sellerID, status).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

var _ domain.OrderRepository = &orderRepositoryImpl{}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepositoryImpl{
		db,
	}
}
