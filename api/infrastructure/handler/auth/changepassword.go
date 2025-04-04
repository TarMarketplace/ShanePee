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

type ChangePasswordBody struct {
	OldPassword string `json:"old_password" example:"VerySecurePassword"`
	NewPassword string `json:"new_password" example:"VerySecureNewPassword"`
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
		Description: "Change password for authenticated user",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *ChangePasswordInput) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		err := h.authSvc.ChangePassword(ctx, *userID, i.Body.OldPassword, i.Body.NewPassword)
		if err != nil {
			if errors.Is(err, service.ErrInvalidOldPassword) {
				return nil, handler.ErrIncorrectOldPassword
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return nil, nil
	})
}
