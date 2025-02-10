package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type ArtToyService interface {
	CreateArtToy(ctx context.Context, artToy *domain.ArtToy) apperror.AppError
	UpdateArtToy(ctx context.Context, id int64, updateBody *domain.ArtToyUpdateBody) apperror.AppError
}

type artToyServiceImpl struct {
	artToyRepo domain.ArtToyRepository
}

func NewArtToyService(artToyRepo domain.ArtToyRepository) ArtToyService {
	return &artToyServiceImpl{artToyRepo: artToyRepo}
}

var _ ArtToyService = &artToyServiceImpl{}

func (svc *artToyServiceImpl) CreateArtToy(ctx context.Context, artToy *domain.ArtToy) apperror.AppError {
	err := svc.artToyRepo.CreateArtToy(ctx, artToy)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (svc *artToyServiceImpl) UpdateArtToy(ctx context.Context, id int64, updateBody *domain.ArtToyUpdateBody) apperror.AppError {
	artToyData := map[string]interface{}{}

	if updateBody.Name != nil {
		artToyData["name"] = *updateBody.Name
	}
	if updateBody.Description != nil {
		artToyData["description"] = *updateBody.Description
	}
	if updateBody.Price != nil {
		artToyData["price"] = *updateBody.Price
	}
	if updateBody.Photo != nil {
		artToyData["photo"] = *updateBody.Photo
	}
	if updateBody.Availability != nil {
		artToyData["availability"] = *updateBody.Availability
	}

	err := svc.artToyRepo.UpdateArtToy(ctx, id, artToyData)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}
