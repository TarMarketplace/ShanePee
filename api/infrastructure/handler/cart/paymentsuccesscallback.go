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

type PaymentSuccessCallbackInput struct {
	Signature string `header:"Stripe-Signature"`
	RawBody   []byte
}

func (h *CartHandler) RegisterPaymentSuccessCallback(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "payment-success-callback",
		Method:      http.MethodPost,
		Path:        "/v1/cart/payment-success-callback",
		Tags:        []string{"Cart"},
		Summary:     "Callback after stripe payment success",
		Description: "Callback after stripe payment success",
	}, func(ctx context.Context, i *PaymentSuccessCallbackInput) (*struct{}, error) {
		err := h.stripeSvc.PaymentSuccessCallback(ctx, i.RawBody, i.Signature)
		if err != nil {
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return nil, nil
	})
}
