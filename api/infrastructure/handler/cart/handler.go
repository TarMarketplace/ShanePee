package cart

import (
	"shanepee.com/api/service"
)

type CartHandler struct {
	cartSvc service.CartService
}

func NewHandler(cartSvc service.CartService) CartHandler {
	return CartHandler{
		cartSvc,
	}
}
