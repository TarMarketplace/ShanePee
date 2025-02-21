package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type ReviewCreateBody struct {
	Rating   int    `json:"rating" example:"5"`
	Comment  string `json:"comment" example:"This is a great art toy"`
	ArtToyID int64  `json:"art_toy_id" example:"97"`
}

type CreateReviewInput struct {
	Body ReviewCreateBody
}

type CreateReviewOutput struct {
	Body *domain.Review
}

func (h *ArtToyHandler) RegisterCreateReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-review",
		Method:      http.MethodPost,
		Path:        "/v1/art-toy/review",
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
		review, err := h.artToySvc.CreateReview(ctx, i.Body.Rating, i.Body.Comment, i.Body.ArtToyID)
		if err != nil {
			if err == service.ErrArtToyNotFound {
				return nil, handler.ErrArtToyNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &CreateReviewOutput{
			Body: review,
		}, nil
	})
}
