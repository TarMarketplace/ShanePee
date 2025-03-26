package review

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetReviewsOfSellerInput struct {
	SellerID int64 `path:"sellerID"`
}

type GetReviewsOfSellerOutput struct {
	Body handler.ArrayResponse[domain.ReviewWithTruncatedBuyer]
}

func (h *ReviewHandler) RegisterGetReviewsOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-reviews-of-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/review/{sellerID}",
		Tags:        []string{"Review"},
		Summary:     "Get Order Reviews of seller",
		Description: "Get order reviews of seller",
	}, func(ctx context.Context, i *GetReviewsOfSellerInput) (*GetReviewsOfSellerOutput, error) {
		data, err := h.reviewSvc.GetReviewsWithTruncatedBuyerBySellerID(ctx, i.SellerID)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return &GetReviewsOfSellerOutput{
			Body: handler.ArrayResponse[domain.ReviewWithTruncatedBuyer]{
				Data: data,
			},
		}, nil
	})
}
