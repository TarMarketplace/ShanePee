package domain

import (
	"context"
	"time"
)

type ChatRepository interface {
	FindChatByID(ctx context.Context, chatID int64) (*ChatMessage, error)
	FindLatestChatsBySenderIDAndReceiverID(ctx context.Context, senderID int64, receiverID int64, latestChatTime time.Time) ([]*ChatMessage, error)
	FindChatListByUserID(ctx context.Context, userID int64, latestChatTime time.Time) ([]*ChatList, error)
	CreateChat(ctx context.Context, chat *ChatMessage) error
}
