package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type UserService interface {
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

func (s *userServiceImpl) UpdateUser(ctx context.Context, id int64, user map[string]interface{}) apperror.AppError {
	err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}
