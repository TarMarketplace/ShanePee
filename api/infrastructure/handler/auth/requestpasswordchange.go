package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

type RequestPasswordChangeBody struct {
	Email string `json:"email" example:"johndoe@example.com"`
}

type CreateRequestPasswordChangeInput struct {
	Body RequestPasswordChangeBody
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
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
