package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type ResetPasswordBody struct {
	RequestID   int64  `json:"request_id"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordInput struct {
	Body ResetPasswordBody
}

func (h *AuthHandler) RegisterResetPassword(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "reset-password",
		Method:      http.MethodPost,
		Path:        "/v1/auth/reset-password",
		Tags:        []string{"Authentication"},
		Summary:     "Reset password",
		Description: "Reset password of a user using token and request id",
	}, func(ctx context.Context, i *ResetPasswordInput) (*struct{}, error) {
		if err := h.authSvc.ResetPassword(ctx, i.Body.RequestID, i.Body.Token, i.Body.NewPassword); err != nil {
			if errors.Is(err, service.ErrInvalidToken) {
				return nil, handler.ErrInvalidToken
			}
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
