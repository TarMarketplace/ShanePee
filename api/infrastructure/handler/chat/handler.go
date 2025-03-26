package chat

import (
	"shanepee.com/api/service"
)

type ChatHandler struct {
	chatSvc service.ChatService
}

func NewHandler(chatSvc service.ChatService) ChatHandler {
	return ChatHandler{
		chatSvc,
	}
}
