package domain

import (
	"context"
	"errors"
)

var ErrArtToyNotFound error = errors.New("art toy not found")

type ArtToyRepository interface {
	FindArtToys(ctx context.Context) ([]*ArtToy, error)
	FindArtToyById(ctx context.Context, id int64) (*ArtToy, error)
}
