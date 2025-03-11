package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrOrderNotFound error = domain.ErrOrderNotFound
)

type OrderService interface {
	GetOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*domain.Order, error)
	GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*domain.OrderWithArtToys, error)
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

func (s *orderServiceImpl) GetOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*domain.Order, error) {
	orders, err := s.orderRepo.FindOrdersByStatus(ctx, status, sellerID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderServiceImpl) GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*domain.OrderWithArtToys, error) {
	return s.orderRepo.FindOrdersWithArtToysBySellerID(ctx, sellerID)
}
