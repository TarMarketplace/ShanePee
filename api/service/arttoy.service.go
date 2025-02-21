package service

import (
	"context"
	"errors"

	"shanepee.com/api/domain"
)

var (
	ErrArtToyNotFound error = domain.ErrArtToyNotFound
	ErrReviewNotFound error = domain.ErrReviewNotFound
	ErrUnauthorized   error = errors.New("unauthorized access")
)

type ArtToyService interface {
	CreateArtToy(ctx context.Context, name string, description string, price float64, photo *string, ownerID int64) (*domain.ArtToy, error)
	UpdateArtToy(ctx context.Context, id int64, updateBody map[string]any, ownerID int64) (*domain.ArtToy, error)
	GetArtToys(ctx context.Context) ([]*domain.ArtToy, error)
	GetArtToyByID(ctx context.Context, id int64) (*domain.ArtToy, error)
	DeleteArtToy(ctx context.Context, id int64, ownerID int64) error
	CreateReview(ctx context.Context, rating int, comment string, artToyID int64) (*domain.Review, error)
	GetReviews(ctx context.Context, artToyID int64) ([]*domain.Review, error)
	UpdateReview(ctx context.Context, id int64, updateBody map[string]any) (*domain.Review, error)
	DeleteReview(ctx context.Context, id int64) error
}

type artToyServiceImpl struct {
	artToyRepo domain.ArtToyRepository
}

func NewArtToyService(artToyRepo domain.ArtToyRepository) ArtToyService {
	return &artToyServiceImpl{
		artToyRepo,
	}
}

var _ ArtToyService = &artToyServiceImpl{}

func (s *artToyServiceImpl) CreateArtToy(ctx context.Context, name string, description string, price float64, photo *string, ownerID int64) (*domain.ArtToy, error) {
	artToy := domain.NewArtToy(name, description, price, photo, ownerID)
	err := s.artToyRepo.CreateArtToy(ctx, artToy)
	if err != nil {
		return nil, err
	}
	return artToy, nil
}

func (s *artToyServiceImpl) UpdateArtToy(ctx context.Context, id int64, updateBody map[string]any, ownerID int64) (*domain.ArtToy, error) {
	artToy, err := s.artToyRepo.FindArtToyByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if artToy.OwnerID != ownerID {
		return nil, ErrUnauthorized
	}
	err = s.artToyRepo.UpdateArtToy(ctx, id, updateBody)
	if err != nil {
		return nil, err
	}
	updatedArtToy, err := s.artToyRepo.FindArtToyByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedArtToy, nil
}

func (s *artToyServiceImpl) GetArtToys(ctx context.Context) ([]*domain.ArtToy, error) {
	artToys, err := s.artToyRepo.FindArtToys(ctx)
	if err != nil {
		return nil, err
	}
	return artToys, nil
}

func (s *artToyServiceImpl) GetArtToyByID(ctx context.Context, id int64) (*domain.ArtToy, error) {
	artToy, err := s.artToyRepo.FindArtToyByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return artToy, nil
}

func (s *artToyServiceImpl) DeleteArtToy(ctx context.Context, id int64, ownerID int64) error {
	artToy, err := s.artToyRepo.FindArtToyByID(ctx, id)
	if err != nil {
		return err
	}
	if artToy.OwnerID != ownerID {
		return ErrUnauthorized
	}
	return s.artToyRepo.DeleteArtToy(ctx, id)
}

func (s *artToyServiceImpl) CreateReview(ctx context.Context, rating int, comment string, artToyID int64) (*domain.Review, error) {
	review := domain.NewReview(rating, comment, artToyID)
	err := s.artToyRepo.CreateReview(ctx, review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *artToyServiceImpl) GetReviews(ctx context.Context, artToyID int64) ([]*domain.Review, error) {
	reviews, err := s.artToyRepo.FindReviewsByArtToyID(ctx, artToyID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *artToyServiceImpl) UpdateReview(ctx context.Context, id int64, updateBody map[string]any) (*domain.Review, error) {
	err := s.artToyRepo.UpdateReview(ctx, id, updateBody)
	if err != nil {
		return nil, err
	}
	updatedReview, err := s.artToyRepo.FindReviewByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return updatedReview, nil
}

func (s *artToyServiceImpl) DeleteReview(ctx context.Context, id int64) error {
	return s.artToyRepo.DeleteReview(ctx, id)
}
