package repository

import (
	"context"

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

func NewArtToyRepository(db *gorm.DB) domain.ArtToyRepository {
	return &artToyRepositoryImpl{
		db: db,
	}
}
