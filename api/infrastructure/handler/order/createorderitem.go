package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type OrderItemCreateBody struct {
	ArtToyID int64 `json:"art_toy_id" example:"97"`
	OrderID  int64 `json:"order_id" example:"97"`
}

type CreateOrderItemInput struct {
	Body OrderItemCreateBody
}

type CreateOrderItemOutput struct {
	Body *domain.OrderItem
}

func (h *OrderHandler) RegisterCreateOrderItem(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-order",
		Method:      http.MethodPost,
		Path:        "/v1/order-item",
		Tags:        []string{"Order"},
		Summary:     "Create Order Item",
		Description: "Create a new order item record",
	}, func(ctx context.Context, i *CreateOrderItemInput) (*CreateOrderItemOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		orderItem, err := h.orderSvc.CreateOrderItem(ctx, i.Body.ArtToyID, i.Body.OrderID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &CreateOrderItemOutput{
			Body: orderItem,
		}, nil
	})
}
