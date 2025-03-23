package domain

import (
	"context"
)

type CartRepository interface {
	AddItemToCart(ctx context.Context, cartItem *CartItem) error
	RemoveItemFromCart(ctx context.Context, ownerID int64, ID int64) error
	ClearItemsByOwnerID(ctx context.Context, ownerID int64) error
	GetCartWithItemByOwnerID(ctx context.Context, ownerID int64) ([]*CartItem, error)
}
