package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// CreateUser implements domain.UserRepository.
func (a *userRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	return a.db.Create(user).Error
}

// UpdateUser implements domain.UserRepository.
func (u *userRepositoryImpl) UpdateUser(ctx context.Context, id int64, user map[string]interface{}) error {
	if err := u.db.Model(domain.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	if err := u.db.Take(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) CreatePasswordChangeRequest(ctx context.Context, passwordChangeRequest *domain.PasswordChangeRequest) error {
	return u.db.Create(passwordChangeRequest).Error
}

func (u *userRepositoryImpl) FindPasswordChangeRequestWithUserByID(ctx context.Context, id int64) (*domain.PasswordChangeRequest, error) {
	var passwordChangeRequest domain.PasswordChangeRequest
	if err := u.db.Joins("User").Take(&passwordChangeRequest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &passwordChangeRequest, nil
}

var _ domain.UserRepository = &userRepositoryImpl{}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db,
	}
}
