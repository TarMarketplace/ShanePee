package arttoy

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetArtToyByIDInput struct {
	ID int64 `path:"id"`
}

type GetArtToyByIDOutput struct {
	Body *domain.ArtToy
}

func (h *ArtToyHandler) RegisterGetArtToyByID(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-art-toy-by-id",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy by ID",
		Description: "Get art toy by id",
	}, func(ctx context.Context, i *GetArtToyByIDInput) (*GetArtToyByIDOutput, error) {
		data, err := h.artToySvc.GetArtToyByID(ctx, int64(i.ID))
		if err != nil {
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetArtToyByIDOutput{
			Body: data,
		}, nil
	})
}
