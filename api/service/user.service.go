package service

import (
	"context"
	"errors"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type UserService interface {
	UpdateUser(ctx context.Context, id int64, body map[string]interface{}) apperror.AppError
	GetUserByID(ctx context.Context, id int64) (*domain.User, apperror.AppError)
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &userServiceImpl{
		userRepo,
	}
}

type userServiceImpl struct {
	userRepo domain.UserRepository
}

var _ UserService = &userServiceImpl{}

func (svc *userServiceImpl) UpdateUser(ctx context.Context, id int64, user map[string]interface{}) apperror.AppError {
	err := svc.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (svc *userServiceImpl) GetUserByID(ctx context.Context, id int64) (*domain.User, apperror.AppError) {
	user, err := svc.userRepo.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, apperror.ErrNotFound("user not found")
		}
	}
	return user, nil
}
