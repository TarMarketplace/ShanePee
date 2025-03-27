package chat

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

type GetChatMessageInput struct {
	UserID int64           `path:"userID"`
	As     domain.UserType `query:"as"`
	ChatID int64           `query:"chatID"`
	Poll   bool            `query:"poll"`
}

type GetChatMessageOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterGetChatMessage(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-chat-message",
		Method:      http.MethodGet,
		Path:        "/v1/chat/{userID}",
		Tags:        []string{"Chat"},
		Summary:     "Get Chat Message",
		Description: "Get chat message. In the chat as a seller or a buyer with the user id, poll message to wait for new message sent by the user id. When receiving messages from the user id or time out, polling again",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *GetChatMessageInput) (*GetChatMessageOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		var data []*domain.ChatMessage
		var err error

		if i.As == domain.Seller {
			data, err = h.chatSvc.GetChatMessageBySeller(ctx, i.UserID, *userID, i.ChatID, i.Poll)
		} else {
			data, err = h.chatSvc.GetChatMessageByBuyer(ctx, *userID, i.UserID, i.ChatID, i.Poll)
		}

		if err != nil {
			if errors.Is(err, service.ErrChatNotBelongToOwner) {
				return nil, handler.ErrChatNotBelongToOwner
			}
			if errors.Is(err, service.ErrChatNotFound) {
				return nil, handler.ErrChatNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &GetChatMessageOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
