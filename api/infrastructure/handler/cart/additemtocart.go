package cart

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type CartItemCreateBody struct {
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
		cartItem, err := h.cartSvc.AddItemToCart(ctx, *userID, i.Body.ArtToyID)
		if err != nil {
			if errors.Is(err, service.ErrArtToyBelongToOwner) {
				return nil, handler.ErrArtToyBelongToOwner
			}
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			if errors.Is(err, service.ErrItemAlreadyAddedToCart) {
				return nil, handler.ErrItemAlreadyAddedToCart
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &AddItemToCartOutput{
			Body: cartItem,
		}, nil
	})
}
