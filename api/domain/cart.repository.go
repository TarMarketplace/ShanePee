package domain

import (
	"context"
)

type CartRepository interface {
	AddArtToyToCart(ctx context.Context, userID int64, artToyID int64) error
	FindArtToysInCart(ctx context.Context, userID int64) ([]*ArtToy, error)
	DeleteArtToysInCart(ctx context.Context, userID int64, artToyIDs []int64) error
}
