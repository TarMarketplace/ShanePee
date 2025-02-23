package handler

import "github.com/danielgtaylor/huma/v2"

var (
	ErrAuthenticationRequired   = huma.Error401Unauthorized("Authentication required")
	ErrForbidden                = huma.Error403Forbidden("Forbidden")
	ErrIntervalServerError      = huma.Error500InternalServerError("Internal server error")
	ErrArtToyNotFound           = huma.Error404NotFound("Art toy not found")
	ErrUserNotFound             = huma.Error404NotFound("User not found")
	ErrCartItemNotFound		    = huma.Error404NotFound("Cart item not found")
	ErrCartItemNotBelongToOwner = huma.Error403Forbidden("Cart item does not belong to the owner")
	ErrIncorrectCredential      = huma.Error403Forbidden("Incorrect email or password")
	ErrInvalidToken             = huma.Error403Forbidden("Invalid token")
	ErrIncorrectOldPassword     = huma.Error401Unauthorized("Incorrect old password")
	ErrUserEmailAlreadyExist    = huma.Error403Forbidden("User with this email already exists")
)

type ArrayResponse[T any] struct {
	Data []*T `json:"data"`
}
