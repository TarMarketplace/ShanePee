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

type CheckoutOutputBody struct {
	URL string `json:"url"`
}

type CheckoutOutput struct {
	Body *CheckoutOutputBody
}

func (h *CartHandler) RegisterCheckout(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "checkout",
		Method:      http.MethodPost,
		Path:        "/v1/cart/checkout",
		Tags:        []string{"Cart"},
		Summary:     "Checkout Items In Cart",
		Description: "Place a new order from items in the cart",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*CheckoutOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		url, err := h.stripeSvc.Checkout(ctx, *userID)
		if err != nil {
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &CheckoutOutput{Body: &CheckoutOutputBody{URL: url}}, nil
	})
}
