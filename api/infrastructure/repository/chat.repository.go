package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type chatRepositoryImpl struct {
	db *gorm.DB
}

func (r *chatRepositoryImpl) CreateChat(ctx context.Context, chat *domain.Chat) error {
	return r.db.Create(chat).Error
}

var _ domain.ChatRepository = &chatRepositoryImpl{}

func NewChatRepository(db *gorm.DB) domain.ChatRepository {
	return &chatRepositoryImpl{
		db,
	}
}
