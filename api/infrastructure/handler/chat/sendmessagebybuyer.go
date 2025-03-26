package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type ChatSentByBuyerCreateBody struct {
	Sender  domain.UserType `json:"sender" enum:"BUYER,SELLER" example:"SELLER"`
	Message string          `json:"message" example:"Hello world"`
}

type SendMessageByBuyerInput struct {
	SellerID int64 `path:"sellerID"`
	Body     ChatSentByBuyerCreateBody
}

type SendMessageByBuyerOutput struct {
	Body *domain.ChatMessage
}

func (h *ChatHandler) RegisterSendMessageByBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "send-message-by-buyer",
		Method:      http.MethodPost,
		Path:        "/v1/buyer/chat/{sellerID}",
		Tags:        []string{"Chat"},
		Summary:     "Send Message By Buyer",
		Description: "Send message by buyer to seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *SendMessageByBuyerInput) (*SendMessageByBuyerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		chat, err := h.chatSvc.SendMessageByBuyer(ctx, *userID, i.SellerID, i.Body.Sender, i.Body.Message)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &SendMessageByBuyerOutput{
			Body: chat,
		}, nil
	})
}
