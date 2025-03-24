package domain

import "context"

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *Chat) error
}
