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

func (a *ARepositoryImpl) FindOne(ctx context.Context, id int64) (*domain.A, error) {
	var data domain.A
	if err := a.db.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (a *ARepositoryImpl) Create(ctx context.Context, data domain.A) error {
	return a.db.Create(&data).Error
}

func (a *ARepositoryImpl) Update(ctx context.Context, id int64, body map[string]interface{}) error {
	if err := a.db.Model(domain.A{}).Where("id = ?", id).Updates(body).Error; err != nil {
		return err
	}
	return nil
}

func (a *ARepositoryImpl) Delete(ctx context.Context, id int64) error {
	return a.db.Delete(&domain.A{}, id).Error
}

var _ domain.ARepository = &ARepositoryImpl{}

func NewARepository(db *gorm.DB) domain.ARepository {
	return &ARepositoryImpl{
		db,
	}
}
