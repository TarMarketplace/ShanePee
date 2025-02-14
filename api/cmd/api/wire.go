//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/email"
	"shanepee.com/api/infrastructure/handler/arttoy"
	"shanepee.com/api/infrastructure/handler/auth"
	"shanepee.com/api/infrastructure/handler/user"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/infrastructure/session"
	"shanepee.com/api/service"
)

func InitializeApp() (App, error) {
	wire.Build(
		service.NewUserService,
		service.NewArtToyService,
		repository.NewUserRepository,
		repository.NewArtToyRepository,
		auth.NewHandler,
		user.NewHandler,
		arttoy.NewHandler,
		service.NewAuthService,
		repository.NewDB,
		config.LoadConfig,
		NewApp,
		session.NewStore,
		session.NewOptions,
		email.NewEmailSender,
	)
	return App{}, nil
}
