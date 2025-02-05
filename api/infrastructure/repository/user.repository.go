package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

var _ domain.UserRepository = &userRepositoryImpl{}

// GetUsers implements domain.UserRepository.
func (u *userRepositoryImpl) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.db.Find(&users).Error
	return users, err
}

// GetUser implements domain.UserRepository.
func (u *userRepositoryImpl) GetUser(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	return user, err
}

// UpdateUser implements domain.UserRepository.
func (u *userRepositoryImpl) UpdateUser(ctx context.Context, user map[string]interface{}) error {
	return u.db.Updates(user).Error
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db,
	}
}
