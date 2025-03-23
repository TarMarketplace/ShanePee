package service

import (
	"context"

	"shanepee.com/api/domain"
)

type UserService interface {
	UpdateUser(ctx context.Context, id int64, body map[string]any) error
	GetSellers(ctx context.Context) ([]*domain.User, error)
	GetSellerByID(ctx context.Context, id int64) (*domain.User, error)
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

func (s *userServiceImpl) UpdateUser(ctx context.Context, id int64, body map[string]any) error {
	if err := s.userRepo.UpdateUser(ctx, id, body); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) GetSellers(ctx context.Context) ([]*domain.User, error) {
	sellers, err := s.userRepo.FindSellers(ctx)
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func (s *userServiceImpl) GetSellerByID(ctx context.Context, id int64) (*domain.User, error) {
	seller, err := s.userRepo.FindSellerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return seller, nil
}
