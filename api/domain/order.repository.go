package domain

import (
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	CreateOrderItems(ctx context.Context, orderItems []*OrderItem) error
	FindOrdersByStatus(ctx context.Context, status string, sellerID int64) ([]*Order, error)
	FindOrdersWithArtToysBySellerID(ctx context.Context, sellerID int64) ([]*Order, error)
	FindOrdersWithArtToysByBuyerID(ctx context.Context, buyerID int64, status string) ([]*Order, error)
	FindOrderByID(ctx context.Context, id int64) (*Order, error)
	UpdateOrder(ctx context.Context, id int64, order map[string]any) error
}
