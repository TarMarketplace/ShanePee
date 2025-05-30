package cart

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetCartOutput struct {
	Body handler.ArrayResponse[domain.CartItem]
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
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &GetCartOutput{
			Body: handler.ArrayResponse[domain.CartItem]{Data: cart},
		}, nil
	})
}
