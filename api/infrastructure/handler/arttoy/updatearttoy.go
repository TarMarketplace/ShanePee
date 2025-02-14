package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type ArtToyUpdateBody struct {
	Name         *string  `json:"name,omitempty" example:"Tuna"`
	Description  *string  `json:"description,omitempty" example:"Delicious fish"`
	Price        *float64 `json:"price,omitempty" example:"55.55"`
	Photo        *string  `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	Availability *bool    `json:"availability,omitempty" example:"false"`
}

type UpdateArtToyInput struct {
	ID   int64 `path:"id"`
	Body ArtToyUpdateBody
}

type UpdateArtToyOutput struct {
	Body *domain.ArtToy
}

func (h *ArtToyHandler) RegisterUpdateArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "update-art-toy",
		Method:      http.MethodPut,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Update Art toy",
		Description: "Update an existing art toy by ID",
	}, func(ctx context.Context, i *UpdateArtToyInput) (*UpdateArtToyOutput, error) {
		userId := handler.GetUserID(ctx)
		if userId == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		updatedArtToy, err := h.artToySvc.UpdateArtToy(ctx, i.ID, i.Body.ToMap(), *userId)
		if err != nil {
			// TODO: find what can cause error
			return nil, handler.ErrIntervalServerError
		}

		return &UpdateArtToyOutput{
			Body: updatedArtToy,
		}, nil
	})
}

func (b *ArtToyUpdateBody) ToMap() map[string]any {
	result := make(map[string]any)

	if b.Name != nil {
		result["name"] = b.Name
	}
	if b.Description != nil {
		result["description"] = b.Description
	}
	if b.Price != nil {
		result["price"] = b.Price
	}
	if b.Photo != nil {
		result["photo"] = b.Photo
	}
	if b.Availability != nil {
		result["availability"] = b.Availability
	}

	return result
}
