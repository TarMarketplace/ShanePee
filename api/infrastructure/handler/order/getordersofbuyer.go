package order

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetOrdersOfBuyerInput struct {
	Status string `query:"status,omitempty" enum:"PREPARING,DELIVERING,COMPLETED"`
}

type GetOrdersOfBuyerOutput struct {
	Body handler.ArrayResponse[domain.Order]
}

func (h *OrderHandler) RegisterGetOrdersOfBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-orders-of-buyer",
		Method:      http.MethodGet,
		Path:        "/v1/buyer/order",
		Tags:        []string{"Order"},
		Summary:     "Get orders of buyer",
		Description: "Get orders of buyer",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *GetOrdersOfBuyerInput) (*GetOrdersOfBuyerOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.orderSvc.GetOrdersWithArtToysByBuyerID(ctx, *userId, i.Status)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &GetOrdersOfBuyerOutput{
			Body: handler.ArrayResponse[domain.Order]{
				Data: data,
			},
		}, nil
	})
}
