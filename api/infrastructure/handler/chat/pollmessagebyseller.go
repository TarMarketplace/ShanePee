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

type PollMessageBySellerInput struct {
	BuyerID int64 `path:"buyerID"`
	ChatID  int64 `query:"chatID"`
}

type PollMessageBySellerOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterPollMessageBySeller(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "poll-message-by-seller",
		Method:      http.MethodGet,
		Path:        "/v1/seller/chat/poll/{buyerID}",
		Tags:        []string{"Chat"},
		Summary:     "Poll Message By Seller",
		Description: "Poll message by seller. In the chat with buyer, poll message to wait for new message sent by the buyer. When receiving messages from the buyer or time out, polling again",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *PollMessageBySellerInput) (*PollMessageBySellerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.chatSvc.PollMessageBySeller(ctx, i.BuyerID, *userID, i.ChatID)
		if err != nil {
			if errors.Is(err, service.ErrChatNotBelongToOwner) {
				return nil, handler.ErrChatNotBelongToOwner
			}
			if errors.Is(err, service.ErrChatNotFound) {
				return nil, handler.ErrChatNotFound
			}
			logrus.Error(err)
			return nil, handler.ErrIntervalServerError
		}
		return &PollMessageBySellerOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
