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

type CartItemCreateBody struct {
	CartID   int64 `json:"cart_id" example:"97"`
	ArtToyID int64 `json:"art_toy_id" example:"97"`
}

type AddItemToCartInput struct {
	Body CartItemCreateBody
}

type AddItemToCartOutput struct {
	Body *domain.CartItem
}

func (h *CartHandler) RegisterAddItemToCart(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "add-item-to-cart",
		Method:      http.MethodPost,
		Path:        "/v1/cart/add-item",
		Tags:        []string{"Cart"},
		Summary:     "Add Item To Cart",
		Description: "Add an item to the cart",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *AddItemToCartInput) (*AddItemToCartOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		cart, err := h.cartSvc.AddItemToCart(ctx, i.Body.CartID, i.Body.ArtToyID)
		if err != nil {
			if errors.Is(err, service.ErrCartAndArtToyNotFound) {
				return nil, handler.ErrCartAndArtToyNotFound
			} else if errors.Is(err, service.ErrCartNotFound) {
				return nil, handler.ErrCartNotFound
			} else if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &AddItemToCartOutput{
			Body: cart,
		}, nil
	})
}
