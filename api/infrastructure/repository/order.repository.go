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

func (r *orderRepositoryImpl) CreateOrderItems(ctx context.Context, orderItems []*domain.OrderItem) error {
	err := r.db.CreateInBatches(orderItems, len(orderItems)).Error
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return domain.ErrArtToyNotFound
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

func (r *orderRepositoryImpl) FindOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*domain.OrderWithArtToys, error) {
	var orders []*domain.Order
	if err := r.db.Preload("OrderItems.ArtToy").Where("seller_id = ?", sellerID).Find(&orders).Error; err != nil {
		return nil, err
	}

	result := []*domain.OrderWithArtToys{}

	for _, order := range orders {
		orderWithArtToys := &domain.OrderWithArtToys{
			Order: order,
			ArtToys: []*domain.ArtToy{},
		}

		for _, orderItem :=  range order.OrderItems {
			orderWithArtToys.ArtToys = append(orderWithArtToys.ArtToys, &orderItem.ArtToy)
		}
		
		result = append(result, orderWithArtToys)
	}

	return result, nil
}

var _ domain.OrderRepository = &orderRepositoryImpl{}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepositoryImpl{
		db,
	}
}
