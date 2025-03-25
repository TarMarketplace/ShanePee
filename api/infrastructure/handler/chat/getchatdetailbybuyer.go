package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetChatDetailByBuyerInput struct {
	SellerID int64 `path:"sellerID"`
}

type GetChatDetailByBuyerOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterGetChatDetailByBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-chat-detail-by-buyer",
		Method:      http.MethodGet,
		Path:        "/v1/buyer/chat/{sellerID}",
		Tags:        []string{"Chat"},
		Summary:     "Get Chat Detail By Buyer",
		Description: "Get chat detail by buyer",
	}, func(ctx context.Context, i *GetChatDetailByBuyerInput) (*GetChatDetailByBuyerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.chatSvc.GetChatDetail(ctx, *userID, i.SellerID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &GetChatDetailByBuyerOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
