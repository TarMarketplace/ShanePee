package service

import (
	"context"
	"errors"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type ArtToyService interface {
	GetArtToys(ctx context.Context) ([]*domain.ArtToy, apperror.AppError)
	GetArtToyById(ctx context.Context, id int64) (*domain.ArtToy, apperror.AppError)
}

func NewArtToyService(artToyRepo domain.ArtToyRepository) ArtToyService {
	return &artToyServiceImpl{
		artToyRepo,
	}
}

type artToyServiceImpl struct {
	artToyRepo domain.ArtToyRepository
}

var _ ArtToyService = &artToyServiceImpl{}

func (s *artToyServiceImpl) GetArtToys(ctx context.Context) ([]*domain.ArtToy, apperror.AppError) {
	artToys, err := s.artToyRepo.FindArtToys(ctx)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return artToys, nil
}

func (s *artToyServiceImpl) GetArtToyById(ctx context.Context, id int64) (*domain.ArtToy, apperror.AppError) {
	artToy, err := s.artToyRepo.FindArtToyById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrArtToyNotFound) {
			return nil, apperror.ErrNotFound("art toy not found")
		}
		return nil, apperror.ErrInternal(err)
	}
	return artToy, nil
}
