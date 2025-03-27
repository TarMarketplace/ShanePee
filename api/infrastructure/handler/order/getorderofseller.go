package order

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetOrderOfSellerInput struct {
	OrderID int64 `path:"orderID"`
}

type GetOrderOfSellerOutput struct {
	Body *domain.Order
}

func (h *OrderHandler) RegisterGetOrderOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-order-of-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/order/{orderID}",
		Tags:        []string{"Order"},
		Summary:     "Get order detail of seller",
		Description: "Get order detail of seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *GetOrderOfSellerInput) (*GetOrderOfSellerOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		order, err := h.orderSvc.GetOrderWithArtToysBySellerID(ctx, i.OrderID, *userId)
		if err != nil {
			if errors.Is(err, service.ErrOrderNotBelongToOwner) {
				return nil, handler.ErrOrderNotBelongToOwner
			}
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &GetOrderOfSellerOutput{
			Body: order,
		}, nil
	})
}
