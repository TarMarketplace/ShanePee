package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]domain.User, apperror.AppError)
	GetUser(ctx context.Context, id int64) (*domain.User, apperror.AppError)
	UpdateUser(ctx context.Context, id int64, body map[string]interface{}) apperror.AppError
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

func (svc *userServiceImpl) GetUsers(ctx context.Context) ([]domain.User, apperror.AppError) {
	users, err := svc.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return users, nil
}

func (svc *userServiceImpl) GetUser(ctx context.Context, id int64) (*domain.User, apperror.AppError) {
	user, err := svc.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return user, nil
}

func (svc *userServiceImpl) UpdateUser(ctx context.Context, id int64, user map[string]interface{}) apperror.AppError {
	err := svc.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}
