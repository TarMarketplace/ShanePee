package service

import (
	"context"

	"shanepee.com/api/domain"
)

type ChatService interface {
	SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error)
	SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error)
}

type chatServiceImpl struct {
	chatRepo domain.ChatRepository
}

func NewChatService(chatRepo domain.ChatRepository) ChatService {
	return &chatServiceImpl{
		chatRepo,
	}
}

var _ ChatService = &chatServiceImpl{}

func (s *chatServiceImpl) SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error) {
	chat := domain.NewChat(buyerID, sellerID, sender, message)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	// TODO: handle something for long polling
	return chat, nil
}

func (s *chatServiceImpl) SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error) {
	chat := domain.NewChat(buyerID, sellerID, sender, message)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	// TODO: handle something for long polling
	return chat, nil
}
