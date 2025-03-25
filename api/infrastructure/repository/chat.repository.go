package repository

import (
	"context"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type chatRepositoryImpl struct {
	db *gorm.DB
}

func (r *chatRepositoryImpl) FindChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.Chat, error) {
	var chats []*domain.Chat
	err := r.db.Table("chats").
		Where("buyer_id = ? AND seller_id = ?", buyerID, sellerID).
		Order("created_at DESC").
		Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
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
