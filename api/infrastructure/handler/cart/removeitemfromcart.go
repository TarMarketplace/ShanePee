package cart

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

type RemoveItemFromCart struct {
	ID int64 `path:"id" example:"97"`
}

func (h *CartHandler) RegisterRemoveItemFromCart(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "remove-item-from-cart",
		Method:      http.MethodDelete,
		Path:        "/v1/cart/remove-item/{id}",
		Tags:        []string{"Cart"},
		Summary:     "Remove Item From Cart",
		Description: "Remove an item from the cart",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *RemoveItemFromCart) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		err := h.cartSvc.RemoveItemFromCart(ctx, *userID, i.ID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		
		return nil, nil
	})
}
