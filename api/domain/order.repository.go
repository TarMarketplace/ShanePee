package domain

import (
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	CreateOrderItems(ctx context.Context, orderItems []*OrderItem) error
	FindOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*Order, error)
}
