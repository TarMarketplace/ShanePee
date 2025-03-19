package review

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/infrastructure/handler"
)

type GetSellerRatingInput struct {
	SellerID int64 `path:"sellerID"`
}

type GetSellerRatingOutput struct {
	Body *float64
}

func (h *ReviewHandler) RegisterGetSellerRating(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-seller-rating",
		Method:      http.MethodGet,
		Path:        "/v1/seller/{sellerID}/rating",
		Tags:        []string{"Art Toy"},
		Summary:     "Get Seller Rating",
		Description: "Get average rating from all art toys of seller",
	}, func(ctx context.Context, i *GetSellerRatingInput) (*GetSellerRatingOutput, error) {
		rating, err := h.reviewSvc.GetSellerRating(ctx, i.SellerID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetSellerRatingOutput{
			Body: rating,
		}, nil
	})
}
