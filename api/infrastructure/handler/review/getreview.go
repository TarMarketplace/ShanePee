package review

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetReviewInput struct {
	ArtToyID int64 `path:"artToyID"`
}

type GetReviewOutput struct {
	Body *domain.Review
}

func (h *ReviewHandler) RegisterGetReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-review",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/review/{artToyID}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy Review",
		Description: "Get art toy review by art toy ID",
	}, func(ctx context.Context, i *GetReviewInput) (*GetReviewOutput, error) {
		review, err := h.reviewSvc.GetReview(ctx, i.ArtToyID)
		if err != nil {
			if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetReviewOutput{
			Body: review,
		}, nil
	})
}
