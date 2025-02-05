package main

import (
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
)

type App struct {
	authHdr handler.AuthHandler
	userHdr handler.UserHandler
	cfg     config.Config
}

func NewApp(authHdr handler.AuthHandler, userHdr handler.UserHandler, cfg config.Config) App {
	return App{
		authHdr,
		userHdr,
		cfg,
	}
}
