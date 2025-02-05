//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/service"
)

func InitializeApp() (App, error) {
	wire.Build(
		handler.NewAHandler,
		service.NewAService,
		service.NewUserService,
		repository.NewARepository,
		repository.NewUserRepository,
		handler.NewAuthHandler,
		handler.NewUserHandler,
		service.NewAuthService,
		repository.NewDB,
		config.LoadConfig,
		NewApp,
	)
	return App{}, nil
}
