package domain

import (
	"context"
)

type ArtToyRepository interface {
	CreateArtToy(ctx context.Context, artToy *ArtToy) error
	FindArtToys(ctx context.Context) ([]*ArtToy, error)
	FindArtToysByOwnerID(ctx context.Context, ownerID int64) ([]*ArtToy, error)
	FindArtToyByID(ctx context.Context, id int64) (*ArtToy, error)
	UpdateArtToy(ctx context.Context, id int64, artToy map[string]any) error
	UpdateArtToysAvailability(ctx context.Context, artToyIDs []int64, available bool) error
	DeleteArtToy(ctx context.Context, id int64) error
}
