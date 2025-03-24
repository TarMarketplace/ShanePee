package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrReviewNotFound error = domain.ErrReviewNotFound
)

type ReviewService interface {
	CreateReview(ctx context.Context, rating int, comment string, orderID int64, buyerID int64) (*domain.Review, error)
}

type reviewServiceImpl struct {
	orderRepo  domain.OrderRepository
	reviewRepo domain.ReviewRepository
}

func NewReviewService(orderRepo domain.OrderRepository, reviewRepo domain.ReviewRepository) ReviewService {
	return &reviewServiceImpl{
		orderRepo,
		reviewRepo,
	}
}

var _ ReviewService = &reviewServiceImpl{}

func (s *reviewServiceImpl) CreateReview(ctx context.Context, rating int, comment string, orderID int64, buyerID int64) (*domain.Review, error) {
	order, err := s.orderRepo.FindOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.BuyerID != buyerID {
		return nil, ErrOrderNotBelongToOwner
	}

	review := domain.NewReview(rating, comment, orderID)
	if err := s.reviewRepo.CreateReview(ctx, review); err != nil {
		return nil, err
	}
	return review, nil
}
