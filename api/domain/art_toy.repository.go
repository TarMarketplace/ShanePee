package domain

import (
	"context"
)

type ArtToyRepository interface {
	FindArtToys(ctx context.Context) ([]*ArtToy, error)
}
