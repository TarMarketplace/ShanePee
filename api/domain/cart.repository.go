package domain

import (
	"context"
)

type CartRepository interface {
	AddArtToyToCart(ctx context.Context, userId int64, artToyId int64) error
	FindArtToysInCart(ctx context.Context, userId int64) ([]*ArtToy, error)
	DeleteArtToysInCart(ctx context.Context, userId int64, artToyIds []int64) error
}
