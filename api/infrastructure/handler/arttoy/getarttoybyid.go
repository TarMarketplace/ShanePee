package arttoy

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetArtToyByIdInput struct {
	Id int `path:"id"`
}

type GetArtToyByIdOutput struct {
	Body *domain.ArtToy
}

func (h *ArtToyHandler) RegisterGetArtToyById(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-art-toy-by-id",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy by ID",
		Description: "Get art toy by id",
	}, func(ctx context.Context, i *GetArtToyByIdInput) (*GetArtToyByIdOutput, error) {
		data, err := h.artToySvc.GetArtToyById(ctx, int64(i.Id))
		if errors.Is(err, domain.ErrArtToyNotFound) {
			return nil, handler.ErrArtToyNotFound
		} else if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetArtToyByIdOutput{
			Body: data,
		}, nil
	})
}
