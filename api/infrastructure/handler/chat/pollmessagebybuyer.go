package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type PollMessageByBuyerOutput struct {
	Body *domain.Chat
}

func (h *ChatHandler) RegisterPollMessageByBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "poll-message-by-buyer",
		Method:      http.MethodGet,
		Path:        "/v1/buyer/chat/poll",
		Tags:        []string{"Chat"},
		Summary:     "Poll Message By Buyer",
		Description: "Poll message by buyer",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *struct{}) (*PollMessageByBuyerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		chat, err := h.chatSvc.PollMessageByBuyer(ctx, *userID)
		if err != nil {
			return nil, handler.ErrIntervalServerError
		}
		return &PollMessageByBuyerOutput{
			Body: chat,
		}, nil
	})
}
