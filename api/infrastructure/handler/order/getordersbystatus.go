package order

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetOrdersByStatusInput struct {
	Status string `path:"status"`
}

type GetOrdersByStatusOutput struct {
	Body handler.ArrayResponse[domain.Order]
}

func (h *OrderHandler) RegisterGetOrdersByStatus(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-orders-by-status",
		Method:      http.MethodGet,
		Path:        "/v1/order/{status}",
		Tags:        []string{"Order"},
		Summary:     "Get Order by Status",
		Description: "Get order by status",
	}, func(ctx context.Context, i *GetOrdersByStatusInput) (*GetOrdersByStatusOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersByStatus(ctx, i.Status, *userId)

		if err != nil {
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetOrdersByStatusOutput{
			Body: handler.ArrayResponse[domain.Order]{
				Data: data,
			},
		}, nil
	})
}
