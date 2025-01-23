package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type AService interface {
	FindA(ctx context.Context) ([]domain.A, apperror.AppError)
	CreateA(ctx context.Context, body domain.ACreateBody) (*domain.A, apperror.AppError)
}

func NewAService(aRepo domain.ARepository) AService {
	return &aServiceImpl{
		aRepo,
	}
}

type aServiceImpl struct {
	aRepo domain.ARepository
}

var _ AService = &aServiceImpl{}

func (svc *aServiceImpl) FindA(ctx context.Context) ([]domain.A, apperror.AppError) {
	data, err := svc.aRepo.FindMany(ctx)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return data, nil
}

func (svc *aServiceImpl) CreateA(ctx context.Context, body domain.ACreateBody) (*domain.A, apperror.AppError) {
	a := domain.CreateAFromBody(body)
	err := svc.aRepo.Create(ctx, a)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return &a, nil
}
