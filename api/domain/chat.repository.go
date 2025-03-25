package domain

import (
	"context"
	"time"
)

type ChatRepository interface {
	FindChatByID(ctx context.Context, chatID int64) (*ChatMessage, error)
	FindChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64) ([]*ChatMessage, error)
	FindLatestChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64, latestChatTime time.Time) ([]*ChatMessage, error)
	CreateChat(ctx context.Context, chat *ChatMessage) error
}
