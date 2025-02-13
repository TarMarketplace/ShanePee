package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/dto"
	"shanepee.com/api/service"
)

type ArtToyHandler struct {
	artToySvc service.ArtToyService
}

func NewArtToyHandler(artToySvc service.ArtToyService) ArtToyHandler {
	return ArtToyHandler{
		artToySvc,
	}
}

type CreateArtToyInput struct {
	Body dto.ArtToyCreateBody
}

func (h *ArtToyHandler) RegisterCreateArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-art-toy",
		Method:      http.MethodPost,
		Path:        "/v1/art-toy",
		Tags:        []string{"Art toy"},
		Summary:     "Create Art toy",
		Description: "Create a new art toy record",
	}, func(ctx context.Context, i *CreateArtToyInput) (*domain.ArtToy, error) {
		userId := GetUserID(ctx)
		if userId == nil {
			return nil, ErrAuthenticationRequired
		}
		artToy := domain.NewArtToy(i.Body.Name, i.Body.Description, i.Body.Price, i.Body.Photo, *userId)
		err := h.artToySvc.CreateArtToy(ctx, artToy)
		if err != nil {
			return nil, ErrIntervalServerError
		}
		return artToy, nil
	})
}

type UpdateArtToyInput struct {
	ID   int64 `path:"id"`
	Body dto.ArtToyUpdateBody
}

func (h *ArtToyHandler) RegisterUpdateArtToy(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "update-art-toy",
		Method:      http.MethodPut,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Update Art toy",
		Description: "Update an existing art toy by ID",
	}, func(ctx context.Context, i *UpdateArtToyInput) (*domain.ArtToy, error) {
		userId := GetUserID(ctx)
		if userId == nil {
			return nil, ErrAuthenticationRequired
		}

		updatedArtToy, err := h.artToySvc.UpdateArtToy(ctx, i.ID, i.Body.ToMap(), *userId)
		if err != nil {
			// TODO: find what can cause error
			return nil, ErrIntervalServerError
		}

		return updatedArtToy, nil
	})
}

type GetArtToysOutput struct {
	Body dto.ArrayResponse[domain.ArtToy]
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
			return nil, ErrIntervalServerError
		}
		return &GetArtToysOutput{
			Body: dto.ArrayResponse[domain.ArtToy]{
				Data: data,
			},
		}, nil
	})
}

type GetArtToyByIdInput struct {
	Id int `path:"id"`
}

func (h *ArtToyHandler) RegisterGetArtToyById(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-art-toy-by-id",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy by ID",
		Description: "Get art toy by id",
	}, func(ctx context.Context, i *GetArtToyByIdInput) (*domain.ArtToy, error) {
		data, err := h.artToySvc.GetArtToyById(ctx, int64(i.Id))
		if errors.Is(err, domain.ErrArtToyNotFound) {
			return nil, huma.Error404NotFound(err.Error())
		}
		return data, nil
	})
}
