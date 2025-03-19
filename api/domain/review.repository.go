package domain

import (
	"context"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *Review) error
	FindReviewByArtToyID(ctx context.Context, artToyID int64) (*Review, error)
	FindReviewerByArtToyID(ctx context.Context, artToyID int64) (*int64, error)
	FindReviewBySellerID(ctx context.Context, sellerID int64) ([]*Review, error)
	UpdateReview(ctx context.Context, artToyID int64, review map[string]any) error
	DeleteReview(ctx context.Context, artToyID int64) error
}
