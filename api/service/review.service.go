package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrReviewNotFound error = domain.ErrReviewNotFound
)

type ReviewService interface {
	CreateReview(ctx context.Context, rating int, comment string, artToyID int64) (*domain.Review, error)
	GetReview(ctx context.Context, artToyID int64) (*domain.Review, error)
	UpdateReview(ctx context.Context, artToyID int64, updateBody map[string]any) (*domain.Review, error)
	DeleteReview(ctx context.Context, artToyID int64) error
}

type reviewServiceImpl struct {
	reviewRepo domain.ReviewRepository
}

func NewReviewService(reviewRepo domain.ReviewRepository) ReviewService {
	return &reviewServiceImpl{
		reviewRepo,
	}
}

var _ ReviewService = &reviewServiceImpl{}

func (s *reviewServiceImpl) CreateReview(ctx context.Context, rating int, comment string, artToyID int64) (*domain.Review, error) {
	review := domain.NewReview(rating, comment, artToyID)
	err := s.reviewRepo.CreateReview(ctx, review)
	if err != nil {
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

func (s *reviewServiceImpl) UpdateReview(ctx context.Context, artToyID int64, updateBody map[string]any) (*domain.Review, error) {
	err := s.reviewRepo.UpdateReview(ctx, artToyID, updateBody)
	if err != nil {
		return nil, err
	}
	updatedReview, err := s.reviewRepo.FindReviewByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	return updatedReview, nil
}

func (s *reviewServiceImpl) DeleteReview(ctx context.Context, artToyID int64) error {
	return s.reviewRepo.DeleteReview(ctx, artToyID)
}
