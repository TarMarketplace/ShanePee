package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

func (h *AuthHandler) RegisterLogout(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "logout",
		Method:      http.MethodPost,
		Path:        "/v1/auth/logout",
		Tags:        []string{"Authentication"},
		Summary:     "Logout User",
		Description: "Logout",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*struct{}, error) {
		session := handler.GetSession(ctx)
		session.Clear()
		newSessionOpts := h.defaultSessionOpts
		newSessionOpts.MaxAge = -1
		session.Options(newSessionOpts)
		if err := session.Save(); err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
