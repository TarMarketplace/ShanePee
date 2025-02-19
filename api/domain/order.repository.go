package domain

import (
	"context"
)

type OrderRepository interface {
	FindOrdersByStatus(ctx context.Context, status string, sellerId int64) ([]*Order, error)
}
