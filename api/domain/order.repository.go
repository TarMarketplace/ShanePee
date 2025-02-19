package domain

import (
	"context"
)

type OrderRepository interface {
	FindOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*Order, error)
}
