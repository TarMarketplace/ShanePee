package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetChatDetailBySellerInput struct {
	BuyerID int64 `path:"buyerID"`
}

type GetChatDetailBySellerOutput struct {
	Body handler.ArrayResponse[domain.Chat]
}

func (h *ChatHandler) RegisterGetChatDetailBySeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-chat-detail-by-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/chat/{buyerID}",
		Tags:        []string{"Chat"},
		Summary:     "Get Chat Detail By Seller",
		Description: "Get chat detail by seller",
	}, func(ctx context.Context, i *GetChatDetailBySellerInput) (*GetChatDetailBySellerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.chatSvc.GetChatDetail(ctx, i.BuyerID, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetChatDetailBySellerOutput{
			Body: handler.ArrayResponse[domain.Chat]{
				Data: data,
			},
		}, nil
	})
}
