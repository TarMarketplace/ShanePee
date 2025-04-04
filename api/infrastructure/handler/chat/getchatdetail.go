package chat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
)

type GetChatDetailInput struct {
	UserID int64           `path:"userID"`
	As     domain.UserType `query:"as"`
}

type GetChatDetailOutput struct {
	Body handler.ArrayResponse[domain.ChatMessage]
}

func (h *ChatHandler) RegisterGetChatDetail(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-chat-detail",
		Method:      http.MethodGet,
		Path:        "/v1/chat/{userID}",
		Tags:        []string{"Chat"},
		Summary:     "Get Chat Detail",
		Description: "Get chat detail as a buyer or a seller with the user id",
	}, func(ctx context.Context, i *GetChatDetailInput) (*GetChatDetailOutput, error) {
		userID := handler.GetUserID(ctx)
		if userID == nil {
			return nil, handler.ErrAuthenticationRequired
		}

		var data []*domain.ChatMessage
		var err error
		if i.As == domain.Seller {
			data, err = h.chatSvc.GetChatDetail(ctx, i.UserID, *userID)
		} else {
			data, err = h.chatSvc.GetChatDetail(ctx, *userID, i.UserID)
		}

		if err != nil {
			return nil, handler.ErrInternalServerError
		}
		return &GetChatDetailOutput{
			Body: handler.ArrayResponse[domain.ChatMessage]{
				Data: data,
			},
		}, nil
	})
}
