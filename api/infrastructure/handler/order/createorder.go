package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type OrderCreateBody struct {
	SellerID int64 `json:"seller_id" example:"97"`
}

type CreateOrderInput struct {
	Body OrderCreateBody
}

type CreateOrderOutput struct {
	Body *domain.Order
}

func (h *OrderHandler) RegisterCreateOrder(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-order",
		Method:      http.MethodPost,
		Path:        "/v1/order",
		Tags:        []string{"Order"},
		Summary:     "Create Order",
		Description: "Create a new order record",
	}, func(ctx context.Context, i *CreateOrderInput) (*CreateOrderOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		order, err := h.orderSvc.CreateOrder(ctx, i.Body.SellerID, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &CreateOrderOutput{
			Body: order,
		}, nil
	})
}
