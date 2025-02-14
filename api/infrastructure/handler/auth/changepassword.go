package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

type ChangePasswordBody struct {
	RequestID   int64  `json:"request_id"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordInput struct {
	Body ChangePasswordBody
}

func (h *AuthHandler) RegisterChangePassword(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "change-password",
		Method:      http.MethodPost,
		Path:        "/v1/auth/change-password",
		Tags:        []string{"Authentication"},
		Summary:     "Change password",
		Description: "Change password of a user using token and request id",
	}, func(ctx context.Context, i *ChangePasswordInput) (*struct{}, error) {
		if err := h.authSvc.ChangePassword(ctx, i.Body.RequestID, i.Body.Token, i.Body.NewPassword); err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
