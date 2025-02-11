package main

import (
	"github.com/gin-contrib/sessions"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
)

type App struct {
	authHdr      handler.AuthHandler
	userHdr      handler.UserHandler
	cfg          config.Config
	sessionStore sessions.Store
}

func NewApp(authHdr handler.AuthHandler, userHdr handler.UserHandler, cfg config.Config, sessionStore sessions.Store) App {
	return App{
		authHdr,
		userHdr,
		cfg,
		sessionStore,
	}
}
