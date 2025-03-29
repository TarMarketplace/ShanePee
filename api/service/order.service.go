package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrOrderNotFound         error = domain.ErrOrderNotFound
	ErrOrderNotBelongToOwner error = domain.ErrOrderNotBelongToOwner
)

type OrderService interface {
	GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64, status string) ([]*domain.Order, error)
	GetOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64, status string) ([]*domain.Order, error)
	GetOrderWithArtToysBySellerID(ctx context.Context, orderID int64, sellerID int64) (*domain.Order, error)
	GetOrderWithArtToysByBuyerID(ctx context.Context, orderID int64, buyerID int64) (*domain.Order, error)
	UpdateOrder(ctx context.Context, id int64, updateBody map[string]any, sellerID int64) (*domain.Order, error)
	CompleteOrder(ctx context.Context, id int64, buyerID int64) (*domain.Order, error)
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

func (s *orderServiceImpl) GetOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64, status string) ([]*domain.Order, error) {
	return s.orderRepo.FindOrdersWithArtToysBySellerID(ctx, sellerID, status)
}

func (s *orderServiceImpl) GetOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64, status string) ([]*domain.Order, error) {
	return s.orderRepo.FindOrdersWithArtToysByBuyerID(ctx, buyerID, status)
}

func (s *orderServiceImpl) GetOrderWithArtToysBySellerID(ctx context.Context, orderID int64, sellerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.SellerID != sellerID {
		return nil, ErrOrderNotBelongToOwner
	}
	return order, nil
}

func (s *orderServiceImpl) GetOrderWithArtToysByBuyerID(ctx context.Context, orderID int64, buyerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.BuyerID != buyerID {
		return nil, ErrOrderNotBelongToOwner
	}
	return order, nil
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, id int64, updateBody map[string]any, sellerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order.SellerID != sellerID {
		return nil, ErrOrderNotBelongToOwner
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

func (s *orderServiceImpl) CompleteOrder(ctx context.Context, id int64, buyerID int64) (*domain.Order, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order.BuyerID != buyerID {
		return nil, ErrOrderNotBelongToOwner
	}

	updateBody := make(map[string]any)
	updateBody["status"] = domain.Completed
	if err = s.orderRepo.UpdateOrder(ctx, id, updateBody); err != nil {
		return nil, err
	}
	updatedOrder, err := s.orderRepo.FindOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
