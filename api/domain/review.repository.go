package domain

import (
	"context"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *Review) error
	FindReviewsBySellerID(ctx context.Context, sellerID int64) ([]*Review, error)
}
