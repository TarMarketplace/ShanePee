package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetMyArtToysOutput struct {
	Body handler.ArrayResponse[domain.ArtToy]
}

func (h *ArtToyHandler) RegisterGetMyArtToys(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-my-art-toys",
		Method:      http.MethodGet,
		Path:        "/v1/my-art-toy",
		Tags:        []string{"Art toy"},
		Summary:     "Get My Art Toys",
		Description: "Get my art toys",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*GetMyArtToysOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.artToySvc.GetMyArtToys(ctx, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetMyArtToysOutput{
			Body: handler.ArrayResponse[domain.ArtToy]{
				Data: data,
			},
		}, nil
	})
}
