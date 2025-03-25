package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type LoginBody struct {
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type LoginInput struct {
	Body LoginBody
}

type LoginOutput struct {
	Body *domain.User
}

func (h *AuthHandler) RegisterLogin(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "login",
		Method:      http.MethodPost,
		Path:        "/v1/auth/login",
		Tags:        []string{"Authentication"},
		Summary:     "Login User",
		Description: "Login",
	}, func(ctx context.Context, i *LoginInput) (*LoginOutput, error) {
		data, err := h.authSvc.Login(ctx, i.Body.Email, i.Body.Password)
		if err != nil {
			if errors.Is(err, service.ErrIncorrectCredential) {
				return nil, handler.ErrIncorrectCredential
			}
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}

		session := handler.GetSession(ctx)
		session.Set(handler.UserIDSessionKey, data.ID)
		if err := session.Save(); err != nil {
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}

		return &LoginOutput{
			Body: data,
		}, nil
	})
}
