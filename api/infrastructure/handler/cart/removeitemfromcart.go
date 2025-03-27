package cart

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
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
			if errors.Is(err, service.ErrCartItemNotFound) {
				return nil, handler.ErrCartItemNotFound
			} else if errors.Is(err, service.ErrCartItemNotBelongToOwner) {
				return nil, handler.ErrCartItemNotBelongToOwner
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return nil, nil
	})
}
