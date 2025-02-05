package service

import (
	"context"

	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type AService interface {
	FindManyA(ctx context.Context) ([]domain.A, apperror.AppError)
	FindOneA(ctx context.Context, id int64) (*domain.A, apperror.AppError)
	CreateA(ctx context.Context, body domain.ACreateBody) (*domain.A, apperror.AppError)
	UpdateA(ctx context.Context, id int64, body map[string]interface{}) apperror.AppError
	DeleteA(ctx context.Context, id int64) apperror.AppError
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

func (svc *aServiceImpl) FindManyA(ctx context.Context) ([]domain.A, apperror.AppError) {
	data, err := svc.aRepo.FindMany(ctx)
	if err != nil {
		return nil, apperror.ErrInternal(err)
	}
	return data, nil
}

func (svc *aServiceImpl) FindOneA(ctx context.Context, id int64) (*domain.A, apperror.AppError) {
	data, err := svc.aRepo.FindOne(ctx, id)
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

func (svc *aServiceImpl) UpdateA(ctx context.Context, id int64, body map[string]interface{}) apperror.AppError {
	err := svc.aRepo.Update(ctx, id, body)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (svc *aServiceImpl) DeleteA(ctx context.Context, id int64) apperror.AppError {
	err := svc.aRepo.Delete(ctx, id)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}
