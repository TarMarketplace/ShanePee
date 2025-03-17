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
	GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*domain.Order, error)
	GetOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64) ([]*domain.Order, error)
	GetSellerOrderWithArtToysByOrderID(ctx context.Context, orderID int64, sellerID int64) (*domain.Order, error)
	UpdateOrder(ctx context.Context, id int64, updateBody map[string]any, sellerID int64) (*domain.Order, error)
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

func (s *orderServiceImpl) GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*domain.Order, error) {
	return s.orderRepo.FindOrdersWithArtToysBySellerID(ctx, sellerID)
}

func (s *orderServiceImpl) GetOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64) ([]*domain.Order, error) {
	return s.orderRepo.FindOrdersWithArtToysByBuyerID(ctx, buyerID)
}

func (s *orderServiceImpl) GetSellerOrderWithArtToysByOrderID(ctx context.Context, orderID int64, sellerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.SellerID != sellerID {
		return nil, ErrUnauthorized
	}
	return order, nil
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, id int64, updateBody map[string]any, sellerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order.SellerID != sellerID {
		return nil, ErrUnauthorized
	}

	if err = s.orderRepo.UpdateOrder(ctx, id, updateBody); err != nil {
		return nil, err
	}
	updatedOrder, err := s.orderRepo.FindOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
