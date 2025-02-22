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

var _ domain.ArtToyRepository = &artToyRepositoryImpl{}

func NewArtToyRepository(db *gorm.DB) domain.ArtToyRepository {
	return &artToyRepositoryImpl{
		db,
	}
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

func (r *artToyRepositoryImpl) FindArtToys(ctx context.Context) ([]*domain.ArtToy, error) {
	var artToys []*domain.ArtToy
	if err := r.db.Find(&artToys).Error; err != nil {
		return nil, err
	}
	return artToys, nil
}

func (r *artToyRepositoryImpl) FindArtToysByOwnerID(ctx context.Context, ownerID int64) ([]*domain.ArtToy, error) {
	var artToys []*domain.ArtToy
	if err := r.db.Where("owner_id = ?", ownerID).Find(&artToys).Error; err != nil {
		return nil, err
	}
	return artToys, nil
}

func (r *artToyRepositoryImpl) FindArtToyByID(ctx context.Context, id int64) (*domain.ArtToy, error) {
	var artToy domain.ArtToy
	if err := r.db.Where("id = ?", id).Take(&artToy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrArtToyNotFound
		}
		return nil, err
	}
	return &artToy, nil
}

func (r *artToyRepositoryImpl) DeleteArtToy(ctx context.Context, id int64) error {
	result := r.db.Delete(&domain.ArtToy{}, id)
	return result.Error
}
