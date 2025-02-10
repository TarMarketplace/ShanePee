package repository

import (
	"context"
	"errors"

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

func (a *artToyRepositoryImpl) FindArtToyById(ctx context.Context, id int64) (*domain.ArtToy, error) {
	var artToy domain.ArtToy
	if err := a.db.Where("id = ?", id).Take(&artToy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrArtToyNotFound
		}
		return nil, err
	}
	return &artToy, nil
}

var _ domain.ArtToyRepository = &artToyRepositoryImpl{}

func NewArtToyRepository(db *gorm.DB) domain.ArtToyRepository {
	return &artToyRepositoryImpl{
		db,
	}
}
