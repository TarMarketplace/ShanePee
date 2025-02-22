package domain

import (
	"context"
)

type CartRepository interface {
	AddItemToCart(ctx context.Context, cartItem *CartItem) error
	GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) (*Cart, error)
}
