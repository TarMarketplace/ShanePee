package arttoy

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type DeleteArtToyInput struct {
	ArtToyID int64 `path:"id" example:"3"`
}

func (h *ArtToyHandler) RegisterDeleteArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "delete-art-toy",
		Method:      http.MethodDelete,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Delete Art Toy",
		Description: "Delete an art toy by ID",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *DeleteArtToyInput) (*struct{}, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		err := h.artToySvc.DeleteArtToy(ctx, i.ArtToyID, *userId)
		if err != nil {
			if errors.Is(err, service.ErrArtToyNotBelongToOwner) {
				return nil, handler.ErrArtToyNotBelongToOwner
			}
			if errors.Is(err, service.ErrArtToyNotFound) {
				return nil, handler.ErrArtToyNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}

		return nil, nil
	})
}
