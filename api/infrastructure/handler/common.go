package handler

import "github.com/danielgtaylor/huma/v2"

var (
	ErrAuthenticationRequired = huma.Error401Unauthorized("Authentication required")
	ErrForbidden              = huma.Error403Forbidden("Forbidden")
	ErrIntervalServerError    = huma.Error500InternalServerError("Internal server error")
	ErrArtToyNotFound         = huma.Error404NotFound("Art toy not found")
	ErrUserNotFound           = huma.Error404NotFound("User not found")
	ErrCartNotFound           = huma.Error404NotFound("Cart not found")
	ErrCartAndArtToyNotFound  = huma.Error404NotFound("Cart and art toy not found")
	ErrOrderNotFound          = huma.Error404NotFound("Order not found")
	ErrOrderAndArtToyNotFound = huma.Error404NotFound("Order and art toy not found")
	ErrIncorrectCredential    = huma.Error403Forbidden("Incorrect email or password")
	ErrInvalidToken           = huma.Error403Forbidden("Invalid token")
	ErrIncorrectOldPassword   = huma.Error401Unauthorized("Incorrect old password")
)

type ArrayResponse[T any] struct {
	Data []*T `json:"data"`
}
