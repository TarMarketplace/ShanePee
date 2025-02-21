package domain

import (
	"context"
)

type ArtToyRepository interface {
	CreateArtToy(ctx context.Context, artToy *ArtToy) error
	UpdateArtToy(ctx context.Context, id int64, artToy map[string]any) error
	FindArtToys(ctx context.Context) ([]*ArtToy, error)
	FindArtToyByID(ctx context.Context, id int64) (*ArtToy, error)
	DeleteArtToy(ctx context.Context, id int64) error
	CreateReview(ctx context.Context, review *Review) error
	FindReviewByID(ctx context.Context, id int64) (*Review, error)
	FindReviewsByArtToyID(ctx context.Context, artToyID int64) ([]*Review, error)
	UpdateReview(ctx context.Context, id int64, review map[string]any) error
	DeleteReview(ctx context.Context, id int64) error
}
