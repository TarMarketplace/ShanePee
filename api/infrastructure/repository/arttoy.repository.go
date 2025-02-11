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

func (r *artToyRepositoryImpl) CreateArtToy(ctx context.Context, artToy *domain.ArtToy) error {
	return r.db.Create(artToy).Error
}

func (r *artToyRepositoryImpl) UpdateArtToy(ctx context.Context, id int64, artToy map[string]interface{}) error {
	var count int64

	if err := r.db.Model(&domain.ArtToy{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrArtToyNotFound
	}
	if err := r.db.Model(&domain.ArtToy{}).Where("id = ?", id).Updates(artToy).Error; err != nil {
		return err
	}
	return nil
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

<<<<<<< HEAD
=======

var _ domain.ArtToyRepository = &artToyRepositoryImpl{}
>>>>>>> parent of 64762c1 (Revert "Merge pull request #17 from TarMarketplace/wwichada26/tar-29-be-api-for-create-art-toys")

func NewArtToyRepository(db *gorm.DB) domain.ArtToyRepository {
	return &artToyRepositoryImpl{
		db: db,
	}
}
