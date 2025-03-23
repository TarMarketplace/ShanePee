package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetSellersOutput struct {
	Body handler.ArrayResponse[domain.User]
}

func (h *UserHandler) RegisterGetSellers(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-sellers",
		Method:      http.MethodGet,
		Path:        "/v1/seller",
		Tags:        []string{"User"},
		Summary:     "Get Sellers",
		Description: "Get sellers",
	}, func(ctx context.Context, i *struct{}) (*GetSellersOutput, error) {
		data, err := h.userSvc.GetSellers(ctx)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetSellersOutput{
			Body: handler.ArrayResponse[domain.User]{
				Data: data,
			},
		}, nil
	})
}
