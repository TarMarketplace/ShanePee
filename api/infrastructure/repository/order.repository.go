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

func (r *orderRepositoryImpl) FindOrdersByStatus(ctx context.Context, status string, sellerId int64) ([]*domain.Order, error) {
	var order []*domain.Order
	if err := r.db.Where("seller_id = ? AND status = ?", sellerId, status).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrderNotFound
		}
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
