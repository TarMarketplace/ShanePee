package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type artToyRepositoryImpl struct {
	db *gorm.DB
}

func (a *artToyRepositoryImpl) FindArtToys(ctx context.Context) ([]*domain.ArtToy, error) {
	var artToys []*domain.ArtToy
	if err := a.db.Find(&artToys).Error; err != nil {
		return nil, err
	}
	return artToys, nil
}

var _ domain.ArtToyRepository = &artToyRepositoryImpl{}

func NewArtToyRepository(db *gorm.DB) domain.ArtToyRepository {
	return &artToyRepositoryImpl{
		db,
	}
}
