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
	OrderID int64 `path:"orderID"`
	Body    ReviewCreateBody
}

type CreateReviewOutput struct {
	Body *domain.Review
}

func (h *ReviewHandler) RegisterCreateReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-review",
		Method:      http.MethodPost,
		Path:        "/v1/order/{orderID}/review",
		Tags:        []string{"Review"},
		Summary:     "Create Review",
		Description: "Create a new review for the seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *CreateReviewInput) (*CreateReviewOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		review, err := h.reviewSvc.CreateReview(ctx, i.Body.Rating, i.Body.Comment, i.OrderID, *userID)
		if err != nil {
			if errors.Is(err, service.ErrOrderNotBelongToOwner) {
				return nil, handler.ErrOrderNotBelongToOwner
			}
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &CreateReviewOutput{
			Body: review,
		}, nil
	})
}
