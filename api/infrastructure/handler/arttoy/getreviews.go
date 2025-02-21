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

type GetReviewsInput struct {
	ArtToyID int64 `path:"art-toy-id"`
}

type GetReviewsOutput struct {
	Body handler.ArrayResponse[domain.Review]
}

func (h *ArtToyHandler) RegisterGetReviews(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-reviews",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/review/{art-toy-id}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy Reviews",
		Description: "Get art toy reviews by art toy id",
	}, func(ctx context.Context, i *GetReviewsInput) (*GetReviewsOutput, error) {
		data, err := h.artToySvc.GetReviews(ctx, i.ArtToyID)
		if err != nil {
			if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetReviewsOutput{
			Body: handler.ArrayResponse[domain.Review]{
				Data: data,
			},
		}, nil
	})
}
