package main

import (
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
)

type App struct {
	aHdr handler.AHandler
	authHdr handler.AuthHandler
	cfg  config.Config
}

func NewApp(aHdr handler.AHandler, authHdr handler.AuthHandler, cfg config.Config) App {
	return App{
		aHdr,
		authHdr,
		cfg,
	}
}
