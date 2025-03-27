package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetArtToysOfSellerInput struct {
	ID int64 `path:"id"`
}

type GetArtToysOfSellerOutput struct {
	Body handler.ArrayResponse[domain.ArtToy]
}

func (h *ArtToyHandler) RegisterGetArtToysOfSeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-art-toys-of-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/{id}/art-toy",
		Tags:        []string{"Art toy"},
		Summary:     "Get Art Toys Of Seller",
		Description: "Get art toys of seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *GetArtToysOfSellerInput) (*GetArtToysOfSellerOutput, error) {
		data, err := h.artToySvc.GetArtToysByOwnerID(ctx, i.ID)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &GetArtToysOfSellerOutput{
			Body: handler.ArrayResponse[domain.ArtToy]{
				Data: data,
			},
		}, nil
	})
}
