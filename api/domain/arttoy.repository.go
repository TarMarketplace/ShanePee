package domain

import (
	"context"
	"errors"
)

var ErrArtToyNotFound error = errors.New("art toy not found")

type ArtToyRepository interface {
	CreateArtToy(ctx context.Context, artToy *ArtToy) error
	UpdateArtToy(ctx context.Context, id int64, artToy map[string]any) error
	FindArtToys(ctx context.Context) ([]*ArtToy, error)
	FindArtToyById(ctx context.Context, id int64) (*ArtToy, error)
}
