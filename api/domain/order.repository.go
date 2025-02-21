package domain

import (
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	CreateOrderItem(ctx context.Context, orderItem *OrderItem) error
	FindOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*Order, error)
}
