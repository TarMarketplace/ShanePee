package domain

import "context"

type ChatRepository interface {
	FindChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64) ([]*Chat, error)
	CreateChat(ctx context.Context, chat *Chat) error
}
