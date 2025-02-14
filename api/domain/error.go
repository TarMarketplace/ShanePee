package domain

import "errors"

var (
	// Known error
	ErrUserNotFound                  error = errors.New("user not found")
	ErrArtToyNotFound                error = errors.New("art toy not found")
	ErrPasswordChangeRequestNotFound error = errors.New("password change request not found")
)
