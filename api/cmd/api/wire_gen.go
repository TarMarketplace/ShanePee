// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/service"
)

// Injectors from wire.go:

func InitializeApp() (App, error) {
	configConfig, err := config.LoadConfig()
	if err != nil {
		return App{}, err
	}
	db, err := repository.NewDB(configConfig)
	if err != nil {
		return App{}, err
	}
	aRepository := repository.NewARepository(db)
	aService := service.NewAService(aRepository)
	aHandler := handler.NewAHandler(aService)
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	app := NewApp(aHandler, authHandler, userHandler, configConfig)
	return app, nil
}
