package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetOrdersByStatusInput struct {
	Status string `path:"status" enum:"PENDING,SHIPPING,COMPLETED"`
}

type GetOrdersByStatusOutput struct {
	Body []*domain.Order
}

func (h *OrderHandler) RegisterGetOrdersByStatus(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-orders-by-status",
		Method:      http.MethodGet,
		Path:        "/v1/order/{status}",
		Tags:        []string{"Order"},
		Summary:     "Get Orders by Status",
		Description: "Get orders by status",
	}, func(ctx context.Context, i *GetOrdersByStatusInput) (*GetOrdersByStatusOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersByStatus(ctx, i.Status, *userId)

		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetOrdersByStatusOutput{
			Body: data,
		}, nil
	})
}
