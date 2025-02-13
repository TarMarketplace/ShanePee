package handler

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-contrib/sessions"
	"shanepee.com/api/domain"
	"shanepee.com/api/dto"
	"shanepee.com/api/service"
)

const userIdSessionKey string = "user_id"

type AuthHandler struct {
	authSvc            service.AuthService
	defaultSessionOpts sessions.Options
}

func NewAuthHandler(authSvc service.AuthService, defaultSessionOpts sessions.Options) AuthHandler {
	return AuthHandler{
		authSvc,
		defaultSessionOpts,
	}
}

type RegisterInput struct {
	Body dto.RegisterBody
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
			return nil, ErrIntervalServerError
		}

		session := GetSession(ctx)
		session.Set(userIdSessionKey, data.ID)
		if err := session.Save(); err != nil {
			return nil, ErrIntervalServerError
		}

		return &RegisterOutput{
			Body: data,
		}, nil
	})
}

type LoginInput struct {
	Body dto.LoginBody
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
			return nil, ErrForbidden
		}

		session := GetSession(ctx)
		session.Set(userIdSessionKey, data.ID)
		if err := session.Save(); err != nil {
			return nil, ErrIntervalServerError
		}

		return &LoginOutput{
			Body: data,
		}, nil
	})
}

func (h *AuthHandler) RegisterLogout(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "logout",
		Method:      http.MethodPost,
		Path:        "/v1/auth/logout",
		Tags:        []string{"Authentication"},
		Summary:     "Logout User",
		Description: "Logout",
	}, func(ctx context.Context, i *LoginInput) (*struct{}, error) {
		session := GetSession(ctx)
		session.Clear()
		newSessionOpts := h.defaultSessionOpts
		newSessionOpts.MaxAge = -1
		session.Options(newSessionOpts)
		if err := session.Save(); err != nil {
			return nil, ErrIntervalServerError
		}
		return nil, nil
	})
}

type CreateRequestPasswordChangeInput struct {
	Body struct {
		Email string `json:"email"`
	}
}

func (h *AuthHandler) RegisterCreatePasswordChangeRequests(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "password-change-requests",
		Method:      http.MethodPost,
		Path:        "/v1/auth/password-change-requests",
		Tags:        []string{"Authentication"},
		Summary:     "Request a password reset",
		Description: "Initiates a password reset process by sending an email with reset instructions",
	}, func(ctx context.Context, i *CreateRequestPasswordChangeInput) (*struct{}, error) {
		if err := h.authSvc.RequestPasswordChange(ctx, i.Body.Email); err != nil {
			return nil, ErrIntervalServerError
		}
		return nil, nil
	})
}

type ChangePasswordInput struct {
	Body dto.ChangePasswordBody
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
			return nil, ErrIntervalServerError
		}
		return nil, nil
	})
}

type GetMeOutput struct {
	Body *domain.User
}

func (h *AuthHandler) RegisterGetMe(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "me",
		Method:      http.MethodPost,
		Path:        "/v1/auth/me",
		Tags:        []string{"Authentication"},
		Summary:     "Get current authenticated user",
		Description: "Get authenticated user from the session",
	}, func(ctx context.Context, i *struct{}) (*GetMeOutput, error) {
		userId := GetUserID(ctx)
		if userId == nil {
			return nil, ErrAuthenticationRequired
		}

		data, err := h.authSvc.GetUserByID(ctx, *userId)
		if err != nil {
			return nil, ErrIntervalServerError
		}
		return &GetMeOutput{
			Body: data,
		}, nil
	})
}
