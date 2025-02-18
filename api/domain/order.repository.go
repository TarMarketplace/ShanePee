package domain

import (
	"context"
)

type OrderRepository interface {
	FindOrdersByStatus(ctx context.Context, status string, userID int64) ([]*Order, error)
	CreateOrder(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, id int64, order map[string]any) error
}
