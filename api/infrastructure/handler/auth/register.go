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

type RegisterBody struct {
	Email    string `json:"email"    example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type RegisterInput struct {
	Body RegisterBody
}

type RegisterOutput struct {
	Body *domain.User
}

func (h *AuthHandler) RegisterRegister(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "register",
		Method:      http.MethodPost,
		Path:        "/v1/auth/register",
		Tags:        []string{"Authentication"},
		Summary:     "Register User",
		Description: "Register",
	}, func(ctx context.Context, i *RegisterInput) (*RegisterOutput, error) {
		data, err := h.authSvc.Register(ctx, i.Body.Email, i.Body.Password)
		if err != nil {
			if errors.Is(err, service.ErrUserEmailAlreadyExist) {
				return nil, handler.ErrUserEmailAlreadyExist
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}

		session := handler.GetSession(ctx)
		session.Set(handler.UserIDSessionKey, data.ID)
		if err := session.Save(); err != nil {
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}

		return &RegisterOutput{
			Body: data,
		}, nil
	})
}

func (b *RegisterBody) IntoUser() *domain.User {
	return domain.NewUser(b.Email, b.Password)
}
