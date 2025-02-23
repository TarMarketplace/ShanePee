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

type ReviewCreateBody struct {
	Rating  int    `json:"rating" example:"5"`
	Comment string `json:"comment" example:"This is a great art toy"`
}

type CreateReviewInput struct {
	ArtToyID int64 `path:"artToyID"`
	Body     ReviewCreateBody
}

type CreateReviewOutput struct {
	Body *domain.Review
}

func (h *ReviewHandler) RegisterCreateReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-review",
		Method:      http.MethodPost,
		Path:        "/v1/art-toy/review/{artToyID}",
		Tags:        []string{"Art toy"},
		Summary:     "Create Art Toy Review",
		Description: "Create a new art toy review record",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *CreateReviewInput) (*CreateReviewOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		review, err := h.reviewSvc.CreateReview(ctx, i.Body.Rating, i.Body.Comment, i.ArtToyID, *userID)
		if err != nil {
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			} else if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &CreateReviewOutput{
			Body: review,
		}, nil
	})
}
