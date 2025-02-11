package service

import (
	"context"
	"errors"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type ArtToyService interface {
	CreateArtToy(ctx context.Context, artToy *domain.ArtToy) apperror.AppError
	UpdateArtToy(ctx context.Context, id int64, updateBody *domain.ArtToyUpdateBody, ownerID int64) (*domain.ArtToy, apperror.AppError)
	GetArtToys(ctx context.Context) ([]*domain.ArtToy, apperror.AppError)
	GetArtToyById(ctx context.Context, id int64) (*domain.ArtToy, apperror.AppError)
}

type artToyServiceImpl struct {
	artToyRepo domain.ArtToyRepository
}

func NewArtToyService(artToyRepo domain.ArtToyRepository) ArtToyService {
	return &artToyServiceImpl{artToyRepo: artToyRepo}
}

var _ ArtToyService = &artToyServiceImpl{}

func (s *artToyServiceImpl) CreateArtToy(ctx context.Context, artToy *domain.ArtToy) apperror.AppError {
	err := s.artToyRepo.CreateArtToy(ctx, artToy)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (s *artToyServiceImpl) UpdateArtToy(ctx context.Context, id int64, updateBody *domain.ArtToyUpdateBody, ownerID int64) (*domain.ArtToy, apperror.AppError) {
	artToyData := map[string]interface{}{
		"name":         updateBody.Name,
		"description":  updateBody.Description,
		"price":        updateBody.Price,
		"availability": updateBody.Availability,
		"owner_id":     ownerID,
	}

	if updateBody.Photo == nil {
		artToyData["photo"] = nil
	} else {
		artToyData["photo"] = *updateBody.Photo
	}

	err := s.artToyRepo.UpdateArtToy(ctx, id, artToyData)
	if err != nil {
		if errors.Is(err, domain.ErrArtToyNotFound) {
			return nil, apperror.ErrNotFound("Art toy not found")
		}
		return nil, apperror.ErrInternal(err)
	}
	updatedArtToy, err := s.artToyRepo.FindArtToyById(ctx, id)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}

	return updatedArtToy, nil
}

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
