package handler

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-contrib/sessions"
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
		var userId int64
		session := ctx.Value(defaultSessionKey).(sessions.Session)
		id := session.Get(userIdSessionKey)
		if id == nil {
			return nil, ErrAuthenticationRequired
		}
		userId = id.(int64)

		updateBody := i.Body.IntoMap()
		err := h.userSvc.UpdateUser(ctx, userId, updateBody)
		if err != nil {
			return nil, ErrIntervalServerError
		}
		return nil, nil
	})
}
