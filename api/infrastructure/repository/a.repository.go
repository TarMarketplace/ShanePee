package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type ARepositoryImpl struct {
	db *gorm.DB
}

func (a *ARepositoryImpl) FindMany(ctx context.Context) ([]domain.A, error) {
	var data []domain.A
	if err := a.db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (a *ARepositoryImpl) Create(ctx context.Context, data domain.A) error {
	return a.db.Create(&data).Error
}

var _ domain.ARepository = &ARepositoryImpl{}

func NewARepository(db *gorm.DB) domain.ARepository {
	return &ARepositoryImpl{
		db,
	}
}
