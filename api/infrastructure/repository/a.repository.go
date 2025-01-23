package repository

import (
	"context"
	"gorm.io/driver/sqlite"
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

// TODO: di database
func NewARepository() domain.ARepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.A{})
	return &ARepositoryImpl{
		db,
	}
}
