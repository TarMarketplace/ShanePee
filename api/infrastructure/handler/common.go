package handler

import (
	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrAuthenticationRequired   = huma.Error401Unauthorized("Authentication required")
	ErrForbidden                = huma.Error403Forbidden("Forbidden")
	ErrInternalServerError      = huma.Error500InternalServerError("Internal server error")
	ErrArtToyNotFound           = huma.Error404NotFound("Art toy not found")
	ErrArtToyBelongToOwner      = huma.Error403Forbidden("Art toy is bought by the owner")
	ErrArtToyNotBelongToOwner   = huma.Error403Forbidden("Art toy does not belong to the owner")
	ErrReviewNotFound           = huma.Error404NotFound("Review not found")
	ErrUserNotFound             = huma.Error404NotFound("User not found")
	ErrOrderNotFound            = huma.Error404NotFound("Order not found")
	ErrOrderNotBelongToOwner    = huma.Error403Forbidden("Order does not belong to the owner")
	ErrCartItemNotFound         = huma.Error404NotFound("Cart item not found")
	ErrCartItemNotBelongToOwner = huma.Error403Forbidden("Cart item does not belong to the owner")
	ErrIncorrectCredential      = huma.Error403Forbidden("Incorrect email or password")
	ErrInvalidToken             = huma.Error403Forbidden("Invalid token")
	ErrIncorrectOldPassword     = huma.Error401Unauthorized("Incorrect old password")
	ErrUserEmailAlreadyExist    = huma.Error403Forbidden("User with this email already exists")
	ErrChatNotFound             = huma.Error404NotFound("Chat not found")
	ErrChatNotBelongToOwner     = huma.Error403Forbidden("Chat does not belong to the owner")
	ErrItemAlreadyAddedToCart   = huma.Error403Forbidden("Item already added to the cart")
	ErrArtToyWasPurchased       = huma.Error403Forbidden("Art toy was purchased")
)

type ArrayResponse[T any] struct {
	Data []*T `json:"data"`
}
