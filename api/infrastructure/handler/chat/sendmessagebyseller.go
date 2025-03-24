package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type ChatSentBySellerCreateBody struct {
	BuyerID int64           `json:"buyer_id" example:"97"`
	Sender  domain.UserType `json:"sender" enum:"BUYER,SELLER" example:"SELLER"`
	Message string          `json:"message" example:"Hello world"`
}

type SendMessageBySellerInput struct {
	Body ChatSentBySellerCreateBody
}

type SendMessageBySellerOutput struct {
	Body *domain.Chat
}

func (h *ChatHandler) RegisterSendMessageBySeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "send-message-by-seller",
		Method:      http.MethodPost,
		Path:        "/v1/seller/chat",
		Tags:        []string{"Chat"},
		Summary:     "Send Message By Seller",
		Description: "Send message by seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *SendMessageBySellerInput) (*SendMessageBySellerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		chat, err := h.chatSvc.SendMessageBySeller(ctx, i.Body.BuyerID, *userID, i.Body.Sender, i.Body.Message)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &SendMessageBySellerOutput{
			Body: chat,
		}, nil
	})
}
