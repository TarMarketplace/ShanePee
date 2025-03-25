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

func (h *CartHandler) RegisterDebugCheckout(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "debug-checkout",
		Method:      http.MethodPost,
		Path:        "/v1/cart/debug-checkout",
		Tags:        []string{"Cart"},
		Summary:     "Checkout in debug mode",
		Description: "Checkout in debug mode, bypass payment step",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		err := h.cartSvc.Checkout(ctx, *userID)
		if err != nil {
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
