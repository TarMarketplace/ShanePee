package handler

import "github.com/danielgtaylor/huma/v2"

const (
	defaultSessionKey = "github.com/gin-contrib/sessions"
)

var (
	ErrAuthenticationRequired = huma.Error401Unauthorized("Authentication required")
	ErrForbidden              = huma.Error403Forbidden("Forbidden")
	ErrIntervalServerError    = huma.Error500InternalServerError("Internal server error")
)
