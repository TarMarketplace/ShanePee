package repository

import (
	"context"

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
	return r.db.Create(orderItem).Error
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
