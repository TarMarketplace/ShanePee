package domain

import "errors"

var (
	// Known error
	ErrUserNotFound                 error = errors.New("user not found")
	ErrArtToyNotFound               error = errors.New("art toy not found")
	ErrReviewNotFound               error = errors.New("review not found")
	ErrCartNotFound                 error = errors.New("cart not found")
	ErrCartAndArtToyNotFound        error = errors.New("cart and art toy not found")
	ErrOrderNotFound                error = errors.New("order not found")
	ErrPasswordResetRequestNotFound error = errors.New("password reset request not found")
)
