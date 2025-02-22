package cart

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetCartOutput struct {
	Body *domain.Cart
}

func (h *CartHandler) RegisterGetCart(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-cart",
		Method:      http.MethodGet,
		Path:        "/v1/cart",
		Tags:        []string{"Cart"},
		Summary:     "Get Cart",
		Description: "Retrieve the user's cart",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*GetCartOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		cart, err := h.cartSvc.GetCartWithItemByOwnerID(ctx, *userID)
		if err != nil {
			if errors.Is(err, service.ErrCartNotFound) {
				return nil, handler.ErrCartNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetCartOutput{
			Body: cart,
		}, nil
	})
}
