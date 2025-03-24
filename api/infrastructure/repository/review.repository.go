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
		return domain.ErrOrderNotFound
	}
	return err
}

func (r *reviewRepositoryImpl) FindReviewsBySellerID(ctx context.Context, sellerID int64) ([]*domain.Review, error) {
	var reviews []*domain.Review
	err := r.db.Table("reviews").
		Joins("JOIN orders ON orders.id = reviews.order_id").
		Where("orders.seller_id = ?", sellerID).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
