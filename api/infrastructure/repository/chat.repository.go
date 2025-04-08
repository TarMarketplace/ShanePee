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

func (r *chatRepositoryImpl) FindLatestChatsBySenderIDAndReceiverID(ctx context.Context, senderID int64, receiverID int64, latestChatTime time.Time) ([]*domain.ChatMessage, error) {
	var chats []*domain.ChatMessage
	err := r.db.Model(&domain.ChatMessage{}).
		Raw("? UNION ? ORDER BY created_at ASC",
			r.db.Model(&domain.ChatMessage{}).
				Where("sender_id = ? AND receiver_id = ? AND created_at > ?", senderID, receiverID, latestChatTime),
			r.db.Model(&domain.ChatMessage{}).
				Where("sender_id = ? AND receiver_id = ? AND created_at > ?", receiverID, senderID, latestChatTime)).
		Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepositoryImpl) FindChatListByUserID(ctx context.Context, userID int64, latestChatTime time.Time) ([]*domain.ChatList, error) {
	var chatList []*domain.ChatList

	queryLatestChat := r.db.Model(&domain.ChatMessage{}).
		Select(`
			CASE 
				WHEN sender_id < receiver_id THEN sender_id
				ELSE receiver_id
			END AS user1, 
			CASE
				WHEN sender_id < receiver_id THEN receiver_id
				ELSE sender_id
			END AS user2,
			MAX(created_at) AS last_chat_time`).
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Group("user1, user2")

	err := r.db.Model(&domain.ChatMessage{}).
		Select("chat_messages.id AS id, users.id AS target_id, users.first_name AS target_first_name, users.last_name AS target_last_name, users.photo AS target_photo, chat_messages.message_type AS last_chat_message_type, chat_messages.content AS last_chat_content, chat_messages.created_at AS last_chat_time").
		Joins(`JOIN (?) AS latest_chat ON 
			(CASE WHEN chat_messages.sender_id < chat_messages.receiver_id THEN chat_messages.sender_id ELSE chat_messages.receiver_id END) = latest_chat.user1 AND 
			(CASE WHEN chat_messages.sender_id < chat_messages.receiver_id THEN chat_messages.receiver_id ELSE chat_messages.sender_id END) = latest_chat.user2 AND
			chat_messages.created_at = latest_chat.last_chat_time AND chat_messages.created_at > ?`, queryLatestChat, latestChatTime).
		Joins("JOIN users ON (chat_messages.sender_id = users.id OR chat_messages.receiver_id = users.id) AND users.id != ?", userID).
		Order("chat_messages.created_at ASC").
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
