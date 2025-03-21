package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetOrdersOfBuyerOutput struct {
	Body handler.ArrayResponse[domain.Order]
}

func (h *OrderHandler) RegisterGetOrdersOfBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-orders-of-buyer",
		Method:      http.MethodGet,
		Path:        "/v1/buyer/orders",
		Tags:        []string{"Order"},
		Summary:     "Get orders of buyer",
		Description: "Get orders of buyer",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*GetOrdersOfBuyerOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersWithArtToysByBuyerID(ctx, *userId)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetOrdersOfBuyerOutput{
			Body: handler.ArrayResponse[domain.Order]{
				Data: data,
			},
		}, nil
	})
}
