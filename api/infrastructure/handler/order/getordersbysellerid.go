package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetOrdersOfSellerOutput struct {
	Body handler.ArrayResponse[domain.OrderWithArtToys]
}

func (h *OrderHandler) RegisterGetOrdersOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-orders-of-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/orders",
		Tags:        []string{"Order"},
		Summary:     "Get orders of seller",
		Description: "Get orders of seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*GetOrdersOfSellerOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersWithArtToysBySellerID(ctx, *userId)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetOrdersOfSellerOutput{
			Body: handler.ArrayResponse[domain.OrderWithArtToys]{
				Data: data,
			},
		}, nil
	})
}
