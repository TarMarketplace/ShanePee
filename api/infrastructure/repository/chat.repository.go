package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type chatRepositoryImpl struct {
	db *gorm.DB
}

func (r *chatRepositoryImpl) FindChatByID(ctx context.Context, chatID int64) (*domain.ChatMessage, error) {
	var chat domain.ChatMessage
	if err := r.db.Take(&chat, chatID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrChatNotFound
		}
		return nil, err
	}
	return &chat, nil
}

func (r *chatRepositoryImpl) FindChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.ChatMessage, error) {
	var chats []*domain.ChatMessage
	err := r.db.Model(&domain.ChatMessage{}).
		Where("buyer_id = ? AND seller_id = ?", buyerID, sellerID).
		Order("created_at ASC").
		Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepositoryImpl) FindLatestChatsByBuyerIDAndSellerID(ctx context.Context, buyerID int64, sellerID int64, latestChatTime time.Time) ([]*domain.ChatMessage, error) {
	var chats []*domain.ChatMessage
	err := r.db.Model(&domain.ChatMessage{}).
		Where("buyer_id = ? AND seller_id = ? AND created_at > ?", buyerID, sellerID, latestChatTime).
		Order("created_at ASC").
		Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepositoryImpl) CreateChat(ctx context.Context, chat *domain.ChatMessage) error {
	return r.db.Create(chat).Error
}

var _ domain.ChatRepository = &chatRepositoryImpl{}

func NewChatRepository(db *gorm.DB) domain.ChatRepository {
	return &chatRepositoryImpl{
		db,
	}
}
