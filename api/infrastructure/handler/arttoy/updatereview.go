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

type ReviewUpdateBody struct {
	Rating  *int    `json:"rating,omitempty" example:"5"`
	Comment *string `json:"comment,omitempty" example:"This is a great art toy"`
}

type UpdateReviewInput struct {
	ID   int64 `path:"id"`
	Body ReviewUpdateBody
}

type UpdateReviewOutput struct {
	Body *domain.Review
}

func (h *ArtToyHandler) RegisterUpdateReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "update-review",
		Method:      http.MethodPatch,
		Path:        "/v1/art-toy/review/{id}",
		Tags:        []string{"Art toy"},
		Summary:     "Update Art Toy Review",
		Description: "Update an existing art toy review by ID",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *UpdateReviewInput) (*UpdateReviewOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.artToySvc.UpdateReview(ctx, i.ID, i.Body.ToMap())
		if err != nil {
			if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &UpdateReviewOutput{
			Body: data,
		}, nil
	})
}

func (b *ReviewUpdateBody) ToMap() map[string]any {
	result := make(map[string]any)

	if b.Rating != nil {
		result["rating"] = b.Rating
	}
	if b.Comment != nil {
		result["comment"] = b.Comment
	}
	return result
}
