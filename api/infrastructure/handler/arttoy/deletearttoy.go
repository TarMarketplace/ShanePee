package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type DeleteArtToyInput struct {
	ArtToyID int64 `path:"id" example:"1"`
}

type DeleteArtToyOutput struct {
	Message string `json:"message" example:"Art toy deleted successfully"`
}

func (h *ArtToyHandler) RegisterDeleteArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "delete-art-toy",
		Method:      http.MethodDelete,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Delete Art Toy",
		Description: "Delete an art toy by ID",
	}, func(ctx context.Context, i *DeleteArtToyInput) (*DeleteArtToyOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		err := h.artToySvc.DeleteArtToy(ctx, i.ArtToyID, *userId)
		if err != nil {
			if err == service.ErrArtToyNotFound {
				return nil, handler.ErrArtToyNotFound
			}
			if err == domain.ErrUnauthorized {
				return nil, handler.ErrForbidden
			}
			return nil, handler.ErrIntervalServerError
		}

		return &DeleteArtToyOutput{
			Message: "Art toy deleted successfully",
		}, nil
	})
}
