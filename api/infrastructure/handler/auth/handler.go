package auth

import (
	"github.com/gin-contrib/sessions"
	"shanepee.com/api/service"
)

type AuthHandler struct {
	authSvc            service.AuthService
	defaultSessionOpts sessions.Options
}

func NewHandler(authSvc service.AuthService, defaultSessionOpts sessions.Options) AuthHandler {
	return AuthHandler{
		authSvc,
		defaultSessionOpts,
	}
}
