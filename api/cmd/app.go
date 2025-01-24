package main

import (
	"shanepee.com/api/infrastructure/handler"
)

type App struct {
	aHdr handler.AHandler
}

func NewApp(aHdr handler.AHandler) App {
	return App{
		aHdr,
	}
}
