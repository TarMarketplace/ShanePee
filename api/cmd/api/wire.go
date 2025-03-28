//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/email"
	"shanepee.com/api/infrastructure/handler/arttoy"
	"shanepee.com/api/infrastructure/handler/auth"
	"shanepee.com/api/infrastructure/handler/cart"
	"shanepee.com/api/infrastructure/handler/chat"
	"shanepee.com/api/infrastructure/handler/order"
	"shanepee.com/api/infrastructure/handler/review"
	"shanepee.com/api/infrastructure/handler/user"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/infrastructure/session"
	"shanepee.com/api/service"
)

func InitializeApp() (App, error) {
	wire.Build(
		service.NewUserService,
		service.NewArtToyService,
		service.NewReviewService,
		service.NewCartService,
		service.NewOrderService,
		service.NewStripeService,
		service.NewChatService,
		repository.NewUserRepository,
		repository.NewArtToyRepository,
		repository.NewReviewRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewChatRepository,
		auth.NewHandler,
		user.NewHandler,
		arttoy.NewHandler,
		review.NewHandler,
		cart.NewHandler,
		order.NewHandler,
		chat.NewHandler,
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
