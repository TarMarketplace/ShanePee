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

type OrderUpdateBody struct {
	TrackingNumber  *string             `json:"tracking_number,omitempty" example:"TH1234567890"`
	DeliveryService *string             `json:"delivery_service,omitempty" example:"Kerry Express"`
	Status          *domain.OrderStatus `json:"status,omitempty" enum:"PREPARING,DELIVERING,COMPLETED" example:"pending"`
}

type UpdateOrderBySellerInput struct {
	ID   int64 `path:"id"`
	Body OrderUpdateBody
}

type UpdateOrderBySellerOutput struct {
	Body *domain.Order
}

func (h *OrderHandler) RegisterUpdateOrderBySeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "update-order-by-seller",
		Method:      http.MethodPatch,
		Path:        "/v1/seller/order/{id}",
		Tags:        []string{"Order"},
		Summary:     "Update Order by Seller",
		Description: "Update tracking number, delivery service, status of an order by seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *UpdateOrderBySellerInput) (*UpdateOrderBySellerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		updatedOrder, err := h.orderSvc.UpdateOrderBySeller(ctx, i.ID, i.Body.ToMap(), *userID)
		if err != nil {
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			}
			return nil, handler.ErrIntervalServerError
		}

		return &UpdateOrderBySellerOutput{
			Body: updatedOrder,
		}, nil
	})
}

func (b *OrderUpdateBody) ToMap() map[string]any {
	result := make(map[string]any)

	if b.TrackingNumber != nil {
		result["tracking_number"] = b.TrackingNumber
	}
	if b.DeliveryService != nil {
		result["delivery_service"] = b.DeliveryService
	}
	if b.Status != nil {
		result["status"] = b.Status
	}
	return result
}
