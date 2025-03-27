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

type PollMessageInput struct {
	UserID int64           `path:"userID"`
	As     domain.UserType `query:"as"`
	ChatID int64           `query:"chatID"`
}

type PollMessageOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterPollMessage(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "poll-message",
		Method:      http.MethodGet,
		Path:        "/v1/chat/poll/{userID}",
		Tags:        []string{"Chat"},
		Summary:     "Poll Message",
		Description: "Poll message. In the chat as a seller or a buyer with the user id, poll message to wait for new message sent by the user id. When receiving messages from the user id or time out, polling again",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *PollMessageInput) (*PollMessageOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		var data []*domain.ChatMessage
		var err error

		if i.As == domain.Seller {
			data, err = h.chatSvc.PollMessageBySeller(ctx, i.UserID, *userID, i.ChatID)
		} else {
			data, err = h.chatSvc.PollMessageByBuyer(ctx, *userID, i.UserID, i.ChatID)
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
		return &PollMessageOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
