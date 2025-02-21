package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrOrderNotFound error = domain.ErrOrderNotFound
)

type OrderService interface {
	CreateOrder(ctx context.Context, sellerID int64, buyerID int64) (*domain.Order, error)
	CreateOrderItem(ctx context.Context, artToyID int64, ownerID int64) (*domain.OrderItem, error)
	GetOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*domain.Order, error)
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

func (s *orderServiceImpl) CreateOrder(ctx context.Context, sellerID int64, buyerID int64) (*domain.Order, error) {
	order := domain.NewOrder(sellerID, buyerID)
	err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderServiceImpl) CreateOrderItem(ctx context.Context, artToyID int64, orderID int64) (*domain.OrderItem, error) {
	orderItem := domain.NewOrderItem(artToyID, orderID)
	err := s.orderRepo.CreateOrderItem(ctx, orderItem)
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (s *orderServiceImpl) GetOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*domain.Order, error) {
	orders, err := s.orderRepo.FindOrdersByStatus(ctx, status, sellerID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
