package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetArtToysOutput struct {
	Body handler.ArrayResponse[domain.ArtToy]
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
		data, err := h.artToySvc.GetArtToys(ctx)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return &GetArtToysOutput{
			Body: handler.ArrayResponse[domain.ArtToy]{
				Data: data,
			},
		}, nil
	})
}
