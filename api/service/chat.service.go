package service

import (
	"context"
	"sync"
	"time"

	"shanepee.com/api/domain"
)

var (
	ErrChatNotFound         error = domain.ErrChatNotFound
	ErrChatNotBelongToOwner error = domain.ErrChatNotBelongToOwner
)

type ChatService interface {
	GetChatDetail(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.Chat, error)
	SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error)
	SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error)
	PollMessageBySeller(ctx context.Context, sellerID int64) (*domain.Chat, error)
	PollMessageByBuyer(ctx context.Context, buyerID int64) (*domain.Chat, error)
}

type chatServiceImpl struct {
	chatRepo          domain.ChatRepository
	subscribedBuyers  map[int64]chan *domain.Chat
	subscribedSellers map[int64]chan *domain.Chat
	sync.Mutex
}

func NewChatService(chatRepo domain.ChatRepository) ChatService {
	return &chatServiceImpl{
		chatRepo:          chatRepo,
		subscribedBuyers:  make(map[int64]chan *domain.Chat),
		subscribedSellers: make(map[int64]chan *domain.Chat),
	}
}

var _ ChatService = &chatServiceImpl{}

func (s *chatServiceImpl) GetChatDetail(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.Chat, error) {
	chats, err := s.chatRepo.FindChatsByBuyerIDAndSellerID(ctx, buyerID, sellerID)
	if err != nil {
		return nil, err
	}
	return chats, err
}

func (s *chatServiceImpl) SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error) {
	chat := domain.NewChat(buyerID, sellerID, sender, message)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedSeller(sellerID, chat)
	return chat, nil
}

func (s *chatServiceImpl) SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, message string) (*domain.Chat, error) {
	chat := domain.NewChat(buyerID, sellerID, sender, message)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedBuyer(buyerID, chat)
	return chat, nil
}

func (s *chatServiceImpl) PollMessageByBuyer(ctx context.Context, buyerID int64) (*domain.Chat, error) {
	s.addSubscribedBuyers(buyerID)
	defer s.removeSubscribedBuyer(buyerID)

	select {
	case newChat := <-s.subscribedBuyers[buyerID]:
		return newChat, nil
	case <-time.After(60 * time.Second):
		return nil, nil
	}
}

func (s *chatServiceImpl) PollMessageBySeller(ctx context.Context, sellerID int64) (*domain.Chat, error) {
	s.addSubscribedSellers(sellerID)
	defer s.removeSubscribedSeller(sellerID)

	select {
	case newChat := <-s.subscribedSellers[sellerID]:
		return newChat, nil
	case <-time.After(60 * time.Second):
		return nil, nil
	}
}

func (s *chatServiceImpl) notifySubscribedBuyer(buyerID int64, chat *domain.Chat) {
	s.Lock()
	defer s.Unlock()

	s.subscribedBuyers[buyerID] <- chat
}

func (s *chatServiceImpl) notifySubscribedSeller(sellerID int64, chat *domain.Chat) {
	s.Lock()
	defer s.Unlock()

	s.subscribedSellers[sellerID] <- chat
}

func (s *chatServiceImpl) addSubscribedBuyers(buyerID int64) {
	s.Lock()
	defer s.Unlock()

	s.subscribedBuyers[buyerID] = make(chan *domain.Chat)
}

func (s *chatServiceImpl) addSubscribedSellers(sellerID int64) {
	s.Lock()
	defer s.Unlock()

	s.subscribedSellers[sellerID] = make(chan *domain.Chat)
}

func (s *chatServiceImpl) removeSubscribedBuyer(buyerID int64) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribedBuyers, buyerID)
}

func (s *chatServiceImpl) removeSubscribedSeller(sellerID int64) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribedSellers, sellerID)
}
