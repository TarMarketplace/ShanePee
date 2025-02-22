package domain

import (
	"context"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart *Cart) error
	AddItemToCart(ctx context.Context, cartItem *CartItem) error
	GetCartByOwnerID(ctx context.Context, ownerID int64) ([]*ArtToy, error)
}
