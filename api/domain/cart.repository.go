package domain

import (
	"context"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart *Cart) error
	AddItemToCart(ctx context.Context, cartItem *CartItem) error
	GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) (*Cart, error)
}
