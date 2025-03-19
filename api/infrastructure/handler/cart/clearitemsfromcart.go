package cart

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

func (h *CartHandler) RegisterClearItemsFromCart(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "clear-items-from-cart",
		Method:      http.MethodDelete,
		Path:        "/v1/cart/clear-items",
		Tags:        []string{"Cart"},
		Summary:     "Clear Items From Cart",
		Description: "Clear items from the cart",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		err := h.cartSvc.ClearItemsByOwnerID(ctx, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
