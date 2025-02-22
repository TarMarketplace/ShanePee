package domain

import (
	"context"
)

type CartRepository interface {
	AddItemToCart(ctx context.Context, cartItem *CartItem) error
}
