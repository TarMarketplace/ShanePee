package service

import (
	"context"

	"shanepee.com/api/domain"
)

var (
	ErrArtToyNotFound error = domain.ErrArtToyNotFound
)

type ArtToyService interface {
	CreateArtToy(ctx context.Context, name string, description string, price float64, photo *string, ownerID int64) (*domain.ArtToy, error)
	UpdateArtToy(ctx context.Context, id int64, updateBody map[string]any, ownerID int64) (*domain.ArtToy, error)
	GetArtToys(ctx context.Context) ([]*domain.ArtToy, error)
	GetArtToyByID(ctx context.Context, id int64) (*domain.ArtToy, error)
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
	err := s.artToyRepo.UpdateArtToy(ctx, id, updateBody)
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
