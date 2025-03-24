package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrReviewNotFound error = domain.ErrReviewNotFound
)

type ReviewService interface {
	CreateReview(ctx context.Context, rating int, comment string, artToyID int64, ownerID int64) (*domain.Review, error)
	GetReview(ctx context.Context, artToyID int64) (*domain.Review, error)
	GetReviewsBySellerID(ctx context.Context, sellerID int64) ([]*domain.Review, error)
	UpdateReview(ctx context.Context, artToyID int64, updateBody map[string]any, ownerID int64) (*domain.Review, error)
	DeleteReview(ctx context.Context, artToyID int64, ownerID int64) error
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

func (s *reviewServiceImpl) CreateReview(ctx context.Context, rating int, comment string, artToyID int64, ownerID int64) (*domain.Review, error) {
	buyerID, err := s.reviewRepo.FindReviewerByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	if *buyerID != ownerID {
		return nil, ErrArtToyNotBelongToOwner
	}

	review := domain.NewReview(rating, comment, artToyID)
	if err := s.reviewRepo.CreateReview(ctx, review); err != nil {
		return nil, err
	}
	return review, nil
}

func (s *reviewServiceImpl) GetReview(ctx context.Context, artToyID int64) (*domain.Review, error) {
	review, err := s.reviewRepo.FindReviewByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *reviewServiceImpl) GetReviewsBySellerID(ctx context.Context, sellerID int64) ([]*domain.Review, error) {
	reviews, err := s.reviewRepo.FindReviewBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *reviewServiceImpl) UpdateReview(ctx context.Context, artToyID int64, updateBody map[string]any, ownerID int64) (*domain.Review, error) {
	buyerID, err := s.reviewRepo.FindReviewerByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	if *buyerID != ownerID {
		return nil, ErrArtToyNotBelongToOwner
	}

	if err := s.reviewRepo.UpdateReview(ctx, artToyID, updateBody); err != nil {
		return nil, err
	}
	updatedReview, err := s.reviewRepo.FindReviewByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	return updatedReview, nil
}

func (s *reviewServiceImpl) DeleteReview(ctx context.Context, artToyID int64, ownerID int64) error {
	buyerID, err := s.reviewRepo.FindReviewerByArtToyID(ctx, artToyID)
	if err != nil {
		return err
	}
	if *buyerID != ownerID {
		return ErrArtToyNotBelongToOwner
	}
	return s.reviewRepo.DeleteReview(ctx, artToyID)
}
