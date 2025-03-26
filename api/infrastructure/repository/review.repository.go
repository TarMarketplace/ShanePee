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

func (r *reviewRepositoryImpl) FindReviewsWithTruncatedBuyerBySellerID(ctx context.Context, sellerID int64) ([]*domain.ReviewWithTruncatedBuyer, error) {
	var reviews []*domain.ReviewWithTruncatedBuyer
	err := r.db.Model(&domain.Review{}).
		Select(`*,
			CASE WHEN LENGTH(users.first_name) > 3 THEN CONCAT(SUBSTRING(users.first_name, 1, 3), '***') ELSE users.first_name END AS buyer_truncated_first_name,
			CASE WHEN LENGTH(users.last_name) > 3 THEN CONCAT(SUBSTRING(users.last_name, 1, 3), '***') ELSE users.last_name END AS buyer_truncated_last_name,
			users.photo AS buyer_photo`).
		Joins("JOIN orders ON orders.id = reviews.order_id").
		Joins("JOIN users ON users.id = orders.buyer_id").
		Where("orders.seller_id = ?", sellerID).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
