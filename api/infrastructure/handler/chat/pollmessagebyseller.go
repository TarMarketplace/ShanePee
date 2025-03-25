package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type PollMessageBySellerOutput struct {
	Body *domain.Chat
}

func (h *ChatHandler) RegisterPollMessageBySeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "poll-message-by-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/chat/poll",
		Tags:        []string{"Chat"},
		Summary:     "Poll Message By Seller",
		Description: "Poll message by seller",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*PollMessageBySellerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		chat, err := h.chatSvc.PollMessageBySeller(ctx, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &PollMessageBySellerOutput{
			Body: chat,
		}, nil
	})
}
