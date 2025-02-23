package review

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type DeleteReviewInput struct {
	ArtToyID int64 `path:"artToyID"`
}

func (h *ReviewHandler) RegisterDeleteReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "delete-review",
		Method:      http.MethodDelete,
		Path:        "/v1/art-toy/review/{artToyID}",
		Tags:        []string{"Art toy"},
		Summary:     "Delete Art Toy Review",
		Description: "Delete an existing art toy review by ID",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *DeleteReviewInput) (*struct{}, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		if err := h.reviewSvc.DeleteReview(ctx, i.ArtToyID, *userID); err != nil {
			if errors.Is(err, service.ErrOrderNotFound) {
				return nil, handler.ErrOrderNotFound
			} else if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return nil, nil
	})
}
