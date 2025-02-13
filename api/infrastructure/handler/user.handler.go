package handler

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/dto"
	"shanepee.com/api/service"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) UserHandler {
	return UserHandler{
		userSvc,
	}
}

type UpdateUserInput struct {
	Body dto.UserUpdateBody
}

func (h *UserHandler) UpdateUser(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "update-user",
		Method:      http.MethodPatch,
		Path:        "/v1/user",
		Tags:        []string{"User"},
		Summary:     "Update User",
		Description: "Update user by id",
	}, func(ctx context.Context, i *UpdateUserInput) (*struct{}, error) {
		userId := GetUserID(ctx)
		if userId == nil {
			return nil, ErrAuthenticationRequired
		}

		updateBody := i.Body.IntoMap()
		err := h.userSvc.UpdateUser(ctx, *userId, updateBody)
		if err != nil {
			return nil, ErrIntervalServerError
		}
		return nil, nil
	})
}
