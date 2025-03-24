package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetArtToysOutput struct {
	Body handler.ArrayResponse[domain.ArtToyWithRating]
}

func (h *ArtToyHandler) RegisterGetArtToys(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-art-toys",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toys",
		Description: "Get art toys",
	}, func(ctx context.Context, i *struct{}) (*GetArtToysOutput, error) {
		data, err := h.artToySvc.GetArtToysWithRating(ctx)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetArtToysOutput{
			Body: handler.ArrayResponse[domain.ArtToyWithRating]{
				Data: data,
			},
		}, nil
	})
}
