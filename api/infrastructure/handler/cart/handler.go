package cart

import (
	"shanepee.com/api/service"
)

type CartHandler struct {
	cartSvc   service.CartService
	stripeSvc service.StripeService
}

func NewHandler(cartSvc service.CartService, stripeSvc service.StripeService) CartHandler {
	return CartHandler{
		cartSvc,
		stripeSvc,
	}
}
