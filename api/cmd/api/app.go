package main

import (
	"github.com/gin-contrib/sessions"
	"shanepee.com/api/config"
	"shanepee.com/api/infrastructure/handler/arttoy"
	"shanepee.com/api/infrastructure/handler/auth"
	"shanepee.com/api/infrastructure/handler/cart"
	"shanepee.com/api/infrastructure/handler/chat"
	"shanepee.com/api/infrastructure/handler/order"
	"shanepee.com/api/infrastructure/handler/review"
	"shanepee.com/api/infrastructure/handler/user"
)

type App struct {
	authHdr      auth.AuthHandler
	userHdr      user.UserHandler
	artToyHdr    arttoy.ArtToyHandler
	reviewHdr    review.ReviewHandler
	cartHdr      cart.CartHandler
	orderHdr     order.OrderHandler
	chatHdr      chat.ChatHandler
	cfg          config.Config
	sessionStore sessions.Store
}

func NewApp(authHdr auth.AuthHandler, userHdr user.UserHandler, artToyHdr arttoy.ArtToyHandler, reviewHdr review.ReviewHandler, cartHdr cart.CartHandler, orderHdr order.OrderHandler, chatHdr chat.ChatHandler, cfg config.Config, sessionStore sessions.Store) App {
	return App{
		authHdr,
		userHdr,
		artToyHdr,
		reviewHdr,
		cartHdr,
		orderHdr,
		chatHdr,
		cfg,
		sessionStore,
	}
}
