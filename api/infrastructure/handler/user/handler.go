package user

import (
	"shanepee.com/api/service"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewHandler(userSvc service.UserService) UserHandler {
	return UserHandler{
		userSvc,
	}
}
