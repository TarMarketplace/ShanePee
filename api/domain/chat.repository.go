package domain

import (
	"context"
	"time"
)

type ChatRepository interface {
	FindChatByID(ctx context.Context, chatID int64) (*Chat, error)
	FindChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64) ([]*Chat, error)
	FindLatestChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64, latestChatTime time.Time) ([]*Chat, error)
	CreateChat(ctx context.Context, chat *Chat) error
}
