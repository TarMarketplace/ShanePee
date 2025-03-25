package domain

import (
	"context"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *Review) error
	FindReviewsWithTruncatedBuyerBySellerID(ctx context.Context, sellerID int64) ([]*ReviewWithTruncatedBuyer, error)
}
