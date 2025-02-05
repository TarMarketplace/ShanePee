package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type AuthService interface {
	Register(ctx context.Context, username string, password string) (*domain.User, apperror.AppError)
}

func NewAuthService(userRepo domain.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo,
	}
}

type authServiceImpl struct {
	userRepo domain.UserRepository
}

var _ AuthService = &authServiceImpl{}

func (a *authServiceImpl) Register(ctx context.Context, username string, password string) (*domain.User, apperror.AppError) {
	passwordByte := []byte(password)
	// TODO: salt
	// TODO: validate password
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		// TODO: properly handle this error
		return nil, apperror.ErrInternal(err)
	}
	user := domain.NewUser(username, hashStr)
	err = a.userRepo.CreateUser(ctx, user)
	if err != nil {
		// TODO: properly handle this error
		return nil, apperror.ErrInternal(err)
	}

	return user, nil
}
