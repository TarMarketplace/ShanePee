package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type ArtToyService interface {
	GetArtToys(ctx context.Context) ([]*domain.ArtToy, apperror.AppError)
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
