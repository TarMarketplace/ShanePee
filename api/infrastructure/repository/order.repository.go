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

func (r *orderRepositoryImpl) FindOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64, status string) ([]*domain.Order, error) {
	query := r.db.Preload("OrderItems.ArtToy").Where("seller_id = ?", sellerID)
	if status != "ALL" {
		query = query.Where("status = ?", status)
	}

	var orders []*domain.Order
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepositoryImpl) FindOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64) ([]*domain.Order, error) {
	var orders []*domain.Order
	if err := r.db.Preload("OrderItems.ArtToy").Where("buyer_id = ?", buyerID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepositoryImpl) FindOrderByID(ctx context.Context, id int64) (*domain.Order, error) {
	var order domain.Order
	if err := r.db.Preload("OrderItems.ArtToy").Where("id = ?", id).Take(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderRepositoryImpl) UpdateOrder(ctx context.Context, id int64, order map[string]any) error {
	var count int64

	if err := r.db.Model(&domain.Order{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrOrderNotFound
	}
	if err := r.db.Model(&domain.Order{}).Where("id = ?", id).Updates(order).Error; err != nil {
		return err
	}
	return nil
}

var _ domain.OrderRepository = &orderRepositoryImpl{}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepositoryImpl{
		db,
	}
}
