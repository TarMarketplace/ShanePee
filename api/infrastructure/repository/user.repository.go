package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

var _ domain.AuthRepository = &authRepositoryImpl{}

// CreateUser implements domain.AuthRepository.
func (a *authRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	return a.db.Create(user).Error
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepositoryImpl{
		db,
	}
}
