package main

import (
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
)

type App struct {
	aHdr handler.AHandler
	cfg  config.Config
}

func NewApp(aHdr handler.AHandler, cfg config.Config) App {
	return App{
		aHdr,
		cfg,
	}
}
