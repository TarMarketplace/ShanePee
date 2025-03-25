package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type GetSellerByIDInput struct {
	ID int64 `path:"id"`
}

type GetSellerByIDOutput struct {
	Body *domain.UserWithReview
}

func (h *UserHandler) RegisterGetSellerByID(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-seller-by-id",
		Method:      http.MethodGet,
		Path:        "/v1/seller/{id}",
		Tags:        []string{"User"},
		Summary:     "Get Seller by ID",
		Description: "Get seller by id",
	}, func(ctx context.Context, i *GetSellerByIDInput) (*GetSellerByIDOutput, error) {
		data, err := h.userSvc.GetSellerByID(ctx, i.ID)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return nil, handler.ErrUserNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return &GetSellerByIDOutput{
			Body: data,
		}, nil
	})
}
