package domain

import "errors"

var (
	// Known error
	ErrUserNotFound                 error = errors.New("user not found")
	ErrUserEmailAlreadyExist        error = errors.New("user with this email already exists")
	ErrArtToyNotFound               error = errors.New("art toy not found")
	ErrArtToyNotBelongToOwner       error = errors.New("art toy does not belong to the owner")
	ErrArtToyBelongToOwner          error = errors.New("art toy is bought by the owner")
	ErrReviewNotFound               error = errors.New("review not found")
	ErrOrderNotFound                error = errors.New("order not found")
	ErrOrderNotBelongToOwner        error = errors.New("order does not belong to the owner")
	ErrPasswordResetRequestNotFound error = errors.New("password reset request not found")
	ErrCartItemNotFound             error = errors.New("cart item not found")
	ErrCartItemNotBelongToOwner     error = errors.New("cart item does not belong to the owner")
	ErrChatNotFound                 error = errors.New("chat not found")
	ErrChatNotBelongToOwner         error = errors.New("chat does not belong to the owner")
	ErrItemAlreadyAddedToCart       error = errors.New("item already added to the cart")
)
