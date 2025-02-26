package domain

import "errors"

var (
	// Known error
	ErrUserNotFound                 error = errors.New("user not found")
	ErrUserEmailAlreadyExist        error = errors.New("user with this email already exists")
	ErrArtToyNotFound               error = errors.New("art toy not found")
	ErrReviewNotFound               error = errors.New("review not found")
	ErrUserAndArtToyNotFound        error = errors.New("user and art toy not found")
	ErrOrderNotFound                error = errors.New("order not found")
	ErrOrderAndArtToyNotFound       error = errors.New("order and art toy not found")
	ErrPasswordResetRequestNotFound error = errors.New("password reset request not found")
	ErrCartItemNotFound             error = errors.New("cart item not found")
	ErrCartItemNotBelongToOwner     error = errors.New("cart item does not belong to the owner")
)
