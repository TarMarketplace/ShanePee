package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

type RequestPasswordResetBody struct {
	Email string `json:"email" example:"johndoe@example.com"`
}

type CreateRequestPasswordResetInput struct {
	Body RequestPasswordResetBody
}

func (h *AuthHandler) RegisterCreatePasswordResetRequests(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "password-reset-requests",
		Method:      http.MethodPost,
		Path:        "/v1/auth/password-reset-requests",
		Tags:        []string{"Authentication"},
		Summary:     "Request a password reset",
		Description: "Initiates a password reset process by sending an email with reset instructions",
	}, func(ctx context.Context, i *CreateRequestPasswordResetInput) (*struct{}, error) {
		if err := h.authSvc.RequestPasswordReset(ctx, i.Body.Email); err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
