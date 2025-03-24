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

type GetReviewsOfSellerInput struct {
	SellerID int64 `path:"sellerID"`
}

type GetReviewsOfSellerOutput struct {
	Body handler.ArrayResponse[domain.ReviewResponse]
}

func (h *ReviewHandler) RegisterGetReviewsOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-reviews-of-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/art-toy/review/{sellerID}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy Reviews of Seller",
		Description: "Get art toy reviews of a specific seller",
	}, func(ctx context.Context, i *GetReviewsOfSellerInput) (*GetReviewsOfSellerOutput, error) {
		reviews, err := h.reviewSvc.GetReviewsBySellerID(ctx, i.SellerID)
		if err != nil {
			if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}

		var response []*domain.ReviewResponse
		for _, review := range reviews {
			response = append(response, domain.NewReviewResponse(review))
		}

		return &GetReviewsOfSellerOutput{
			Body: handler.ArrayResponse[domain.ReviewResponse]{Data: response},
		}, nil
	})
}
