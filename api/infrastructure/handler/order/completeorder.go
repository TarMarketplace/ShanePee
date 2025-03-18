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

type CompleteOrderInput struct {
	ID int64 `path:"id"`
}

type CompleteOrderOutput struct {
	Body *domain.Order
}

func (h *OrderHandler) RegisterCompleteOrder(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "complete-order",
		Method:      http.MethodPatch,
		Path:        "/v1/buyer/order/{id}",
		Tags:        []string{"Order"},
		Summary:     "Complete Order by Buyer",
		Description: "Update status to completed of an order by buyer",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *CompleteOrderInput) (*CompleteOrderOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		updatedOrder, err := h.orderSvc.CompleteOrder(ctx, i.ID, *userID)
		if err != nil {
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			}
			return nil, handler.ErrIntervalServerError
		}

		return &CompleteOrderOutput{
			Body: updatedOrder,
		}, nil
	})
}
