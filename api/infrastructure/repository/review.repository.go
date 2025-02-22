package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type reviewRepositoryImpl struct {
	db *gorm.DB
}

var _ domain.ReviewRepository = &reviewRepositoryImpl{}

func NewReviewRepository(db *gorm.DB) domain.ReviewRepository {
	return &reviewRepositoryImpl{
		db,
	}
}

func (r *reviewRepositoryImpl) CreateReview(ctx context.Context, review *domain.Review) error {
	err := r.db.Create(review).Error
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return domain.ErrArtToyNotFound
	}
	return err
}

func (r *reviewRepositoryImpl) FindReviewByArtToyID(ctx context.Context, artToyID int64) (*domain.Review, error) {
	var review domain.Review
	if err := r.db.Where("art_toy_id = ?", artToyID).Take(&review).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrReviewNotFound
		}
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepositoryImpl) UpdateReview(ctx context.Context, artToyID int64, review map[string]interface{}) error {
	var count int64
	if err := r.db.Model(&domain.Review{}).Where("art_toy_id = ?", artToyID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrReviewNotFound
	}
	if err := r.db.Model(&domain.Review{}).Where("art_toy_id = ?", artToyID).Updates(review).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepositoryImpl) DeleteReview(ctx context.Context, artToyID int64) error {
	if err := r.db.Where("art_toy_id = ?", artToyID).Delete(&domain.Review{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrReviewNotFound
		}
		return err
	}
	return nil
}
