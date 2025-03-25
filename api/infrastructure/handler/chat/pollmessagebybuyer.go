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

type PollMessageByBuyerInput struct {
	SellerID int64 `path:"sellerID"`
	ChatID   int64 `query:"chatID"`
}

type PollMessageByBuyerOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterPollMessageByBuyer(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "poll-message-by-buyer",
		Method:      http.MethodGet,
		Path:        "/v1/buyer/chat/poll/{chatID}",
		Tags:        []string{"Chat"},
		Summary:     "Poll Message By Buyer",
		Description: "Poll message by buyer. In the chat with seller, poll message to wait for new message sent by the seller. When receiving messages from the seller or time out, polling again",
		Security: []map[string][]string{
			{"sessionId": {}},
		},
	}, func(ctx context.Context, i *PollMessageByBuyerInput) (*PollMessageByBuyerOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}
		data, err := h.chatSvc.PollMessageByBuyer(ctx, *userID, i.SellerID, i.ChatID)
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
		return &PollMessageByBuyerOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
