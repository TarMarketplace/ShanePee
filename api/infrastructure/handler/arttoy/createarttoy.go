package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type ArtToyCreateBody struct {
	Name        string  `json:"name" example:"Tuna"`
	Description string  `json:"description" example:"Delicious fish"`
	Price       float64 `json:"price" example:"55.55"`
	Photo       *string `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..." nullable:"true"`
}

type CreateArtToyInput struct {
	Body ArtToyCreateBody
}

type CreateArtToyOutput struct {
	Body *domain.ArtToy
}

func (h *ArtToyHandler) RegisterCreateArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-art-toy",
		Method:      http.MethodPost,
		Path:        "/v1/art-toy",
		Tags:        []string{"Art toy"},
		Summary:     "Create Art toy",
		Description: "Create a new art toy record",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *CreateArtToyInput) (*CreateArtToyOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		artToy, err := h.artToySvc.CreateArtToy(ctx, i.Body.Name, i.Body.Description, i.Body.Price, i.Body.Photo, *userID)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &CreateArtToyOutput{
			Body: artToy,
		}, nil
	})
}
