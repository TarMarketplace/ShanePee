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

func (a *userRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	return a.db.Create(user).Error
}

func (u *userRepositoryImpl) UpdateUser(ctx context.Context, id int64, user map[string]any) error {
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

func (u *userRepositoryImpl) CreatePasswordResetRequest(ctx context.Context, passwordResetRequest *domain.PasswordResetRequest) error {
	return u.db.Create(passwordResetRequest).Error
}

func (u *userRepositoryImpl) FindPasswordResetRequestWithUserByID(ctx context.Context, id int64) (*domain.PasswordResetRequest, error) {
	var passwordResetRequest domain.PasswordResetRequest
	if err := u.db.Joins("User").Take(&passwordResetRequest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPasswordResetRequestNotFound
		}
		return nil, err
	}
	return &passwordResetRequest, nil
}

func (u *userRepositoryImpl) UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error {
	user := &domain.User{ID: id}
	return u.db.Model(user).Updates(domain.User{PasswordHash: passwordHash}).Error
}

func (u *userRepositoryImpl) DeletePasswordResetRequestByID(ctx context.Context, id int64) error {
	return u.db.Delete(&domain.PasswordResetRequest{}, id).Error
}

var _ domain.UserRepository = &userRepositoryImpl{}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db,
	}
}
