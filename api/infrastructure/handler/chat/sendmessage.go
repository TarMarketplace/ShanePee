package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type ChatMessageCreateBody struct {
	Content string `json:"content" example:"Hello world"`
}

type SendMessageInput struct {
	UserID int64 `path:"userID"`
	Body   ChatMessageCreateBody
}

type SendMessageOutput struct {
	Body *domain.ChatMessage
}

func (h *ChatHandler) RegisterSendMessage(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "send-message",
		Method:      http.MethodPost,
		Path:        "/v1/chat/send/{userID}",
		Tags:        []string{"Chat"},
		Summary:     "Send Message",
		Description: "Send new message to the user id",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *SendMessageInput) (*SendMessageOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		chat, err := h.chatSvc.SendMessage(ctx, *userID, i.UserID, i.Body.Content)
		if err != nil {
			return nil, handler.ErrInternalServerError
		}
		return &SendMessageOutput{
			Body: chat,
		}, nil
	})
}
