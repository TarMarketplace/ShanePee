package cart

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

func (h *CartHandler) RegisterCheckout(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "checkout",
		Method:      http.MethodPost,
		Path:        "/v1/cart/checkout",
		Tags:        []string{"Cart"},
		Summary:     "Checkout Items In Cart",
		Description: "Place a new order from items in the cart",
	}, func(ctx context.Context, i *struct{}) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		err := h.cartSvc.Checkout(ctx, *userID)
		if err != nil {
			if errors.Is(err, service.ErrOrderAndArtToyNotFound) {
				return nil, handler.ErrOrderAndArtToyNotFound
			} else if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			} else if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
