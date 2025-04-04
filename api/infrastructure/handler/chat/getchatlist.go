package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetChatListInput struct {
	ChatID int64 `query:"chatID"`
	Poll   bool  `query:"poll"`
}

type GetChatListOutput struct {
	Body handler.ArrayResponse[domain.ChatList]
}

func (h *ChatHandler) RegisterGetChatList(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-chat-list",
		Method:      http.MethodGet,
		Path:        "/v1/chatlist",
		Tags:        []string{"Chat"},
		Summary:     "Get Chat List",
		Description: "Get chat list. From all chats with the user id, poll chat to wait for new message sent by users connected with the user id. When receiving messages from the users or time out, polling again",
	}, func(ctx context.Context, i *GetChatListInput) (*GetChatListOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		data, err := h.chatSvc.GetChatList(ctx, *userID, i.Poll, i.ChatID)
		if err != nil {
			return nil, handler.ErrInternalServerError
		}
		return &GetChatListOutput{
			Body: handler.ArrayResponse[domain.ChatList]{
				Data: data,
			},
		}, nil
	})
}
