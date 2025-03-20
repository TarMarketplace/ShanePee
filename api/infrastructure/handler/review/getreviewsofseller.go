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
	Body handler.ArrayResponse[domain.Review]
}

func (h *ReviewHandler) RegisterGetReviewsOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-review",
		Method:      http.MethodGet,
		Path:        "/v1/seller/art-toy/review/{sellerID}",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toy Reviews of seller",
		Description: "Get art toy reviews of seller",
	}, func(ctx context.Context, i *GetReviewsOfSellerInput) (*GetReviewsOfSellerOutput, error) {
		data, err := h.reviewSvc.GetReviewsBySellerID(ctx, i.SellerID)
		if err != nil {
			if errors.Is(err, service.ErrReviewNotFound) {
				return nil, handler.ErrReviewNotFound
			}
			return nil, handler.ErrIntervalServerError
		}
		return &GetReviewsOfSellerOutput{
			Body: handler.ArrayResponse[domain.Review]{
				Data: data,
			},
		}, nil
	})
}
