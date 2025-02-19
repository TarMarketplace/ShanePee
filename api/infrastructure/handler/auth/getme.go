package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetMeOutput struct {
	Body *domain.User
}

func (h *AuthHandler) RegisterGetMe(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "me",
		Method:      http.MethodGet,
		Path:        "/v1/auth/me",
		Tags:        []string{"Authentication"},
		Summary:     "Get current authenticated user",
		Description: "Get authenticated user from the session",
	}, func(ctx context.Context, i *struct{}) (*GetMeOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		data, err := h.authSvc.GetUserByID(ctx, *userID)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return nil, handler.ErrUserNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetMeOutput{
			Body: data,
		}, nil
	})
}
