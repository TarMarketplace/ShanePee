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

func (r *chatRepositoryImpl) FindChatListByUserID(ctx context.Context, userID int64, latestChatTime time.Time) ([]*domain.ChatList, error) {
	var chatList []*domain.ChatList
	err := r.db.Raw("? UNION ?",
		r.db.Model(&domain.ChatMessage{}).
			Select("chat_messages.id AS id, chat_messages.buyer_id AS target_id, chat_messages.sender AS target_type, users.first_name AS target_first_name, users.last_name AS target_last_name, users.photo AS target_photo, chat_messages.content AS last_chat_message, chat_messages.created_at AS last_chat_time").
			Joins(`JOIN (
				SELECT seller_id, MAX(chat_messages.created_at) AS max_created_at
				FROM chat_messages
				GROUP BY seller_id
			) AS latest ON latest.seller_id = chat_messages.seller_id AND chat_messages.created_at = latest.max_created_at`).
			Joins("JOIN users ON users.id = chat_messages.buyer_id").
			Where("chat_messages.buyer_id = ? AND chat_messages.created_at > ?", userID, latestChatTime),
		r.db.Model(&domain.ChatMessage{}).
			Select("chat_messages.id AS id, chat_messages.seller_id AS target_id, chat_messages.sender AS target_type, users.first_name AS target_first_name, users.last_name AS target_last_name, users.photo AS target_photo, chat_messages.content AS last_chat_message, chat_messages.created_at AS last_chat_time").
			Joins(`JOIN (
				SELECT buyer_id, MAX(chat_messages.created_at) AS max_created_at
				FROM chat_messages
				GROUP BY buyer_id
			) AS latest ON latest.buyer_id = chat_messages.buyer_id AND chat_messages.created_at = latest.max_created_at`).
			Joins("JOIN users ON users.id = chat_messages.seller_id").
			Where("chat_messages.seller_id = ? AND chat_messages.created_at > ?", userID, latestChatTime),
	).Order("chat_messaages.created_at ASC").
		Find(&chatList).Error
	if err != nil {
		return nil, err
	}
	return chatList, nil
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
