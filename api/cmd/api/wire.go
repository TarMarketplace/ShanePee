//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/email"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/infrastructure/session"
	"shanepee.com/api/service"
)

func InitializeApp() (App, error) {
	wire.Build(
		service.NewUserService,
		repository.NewUserRepository,
		handler.NewAuthHandler,
		handler.NewUserHandler,
		service.NewAuthService,
		repository.NewDB,
		config.LoadConfig,
		NewApp,
		session.NewStore,
		session.NewOptions,
		email.NewSenderFromConfig,
	)
	return App{}, nil
}
