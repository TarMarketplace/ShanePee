package order

import (
	"shanepee.com/api/service"
)

type OrderHandler struct {
	orderSvc service.OrderService
}

func NewHandler(orderSvc service.OrderService) OrderHandler {
	return OrderHandler{
		orderSvc,
	}
}
