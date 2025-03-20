package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetOrdersOfSellerInput struct {
	Status string `query:"status" enum:"PREPARING,DELIVERING,COMPLETED"`
}

type GetOrdersOfSellerOutput struct {
	Body handler.ArrayResponse[domain.Order]
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
	}, func(ctx context.Context, i *GetOrdersOfSellerInput) (*GetOrdersOfSellerOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersWithArtToysBySellerID(ctx, *userId, i.Status)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetOrdersOfSellerOutput{
			Body: handler.ArrayResponse[domain.Order]{
				Data: data,
			},
		}, nil
	})
}
