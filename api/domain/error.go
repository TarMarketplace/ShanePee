package domain

import "errors"

var (
	// Known error
	ErrUserNotFound                 error = errors.New("user not found")
	ErrArtToyNotFound               error = errors.New("art toy not found")
	ErrPasswordResetRequestNotFound error = errors.New("password reset request not found")
)
