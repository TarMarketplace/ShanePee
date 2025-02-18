package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrOrderNotFound error = domain.ErrOrderNotFound
)

type OrderService interface {
	// CreateOrder(ctx context.Context, buyerId int64, sellerId int64) (*domain.Order, error)
	GetOrdersByStatus(ctx context.Context, status string, id int64) ([]*domain.Order, error)
}

type orderServiceImpl struct {
	orderRepo domain.OrderRepository
}

func NewOrderService(orderRepo domain.OrderRepository) OrderService {
	return &orderServiceImpl{
		orderRepo,
	}
}

var _ OrderService = &orderServiceImpl{}

// func (s *orderServiceImpl) CreateOrder(ctx context.Context, buyerId int64, sellerId int64) (*domain.Order, error) {
// 	order := domain.NewOrder(sellerId, buyerId)
// 	err := s.orderRepo.CreateOrder(ctx, order)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return order, nil
// }

func (s *orderServiceImpl) GetOrdersByStatus(ctx context.Context, status string, id int64) ([]*domain.Order, error) {
	orders, err := s.orderRepo.FindOrdersByStatus(ctx, status, id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
