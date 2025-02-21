package cart

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type CreateCartOutput struct {
	Body *domain.Cart
}

func (h *CartHandler) RegisterCreateCart(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-cart",
		Method:      http.MethodPost,
		Path:        "/v1/cart",
		Tags:        []string{"Cart"},
		Summary:     "Create Cart",
		Description: "Create a new cart record",
	}, func(ctx context.Context, i *struct{}) (*CreateCartOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		cart, err := h.cartSvc.CreateCart(ctx, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &CreateCartOutput{
			Body: cart,
		}, nil
	})
}
