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
	GetChatDetail(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.ChatMessage, error)
	SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error)
	SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error)
	PollMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, chatID int64) ([]*domain.ChatMessage, error)
	PollMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, chatID int64) ([]*domain.ChatMessage, error)
}

type chatServiceImpl struct {
	chatRepo          domain.ChatRepository
	subscribedBuyers  map[int64]map[chan *domain.ChatMessage]struct{}
	subscribedSellers map[int64]map[chan *domain.ChatMessage]struct{}
	sync.Mutex
}

func NewChatService(chatRepo domain.ChatRepository) ChatService {
	return &chatServiceImpl{
		chatRepo:          chatRepo,
		subscribedBuyers:  make(map[int64]map[chan *domain.ChatMessage]struct{}),
		subscribedSellers: make(map[int64]map[chan *domain.ChatMessage]struct{}),
	}
}

var _ ChatService = &chatServiceImpl{}

func (s *chatServiceImpl) GetChatDetail(ctx context.Context, buyerID int64, sellerID int64) ([]*domain.ChatMessage, error) {
	chats, err := s.chatRepo.FindChatsByBuyerIDAndSellerID(ctx, buyerID, sellerID)
	if err != nil {
		return nil, err
	}
	return chats, err
}

func (s *chatServiceImpl) SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error) {
	chat := domain.NewChatMessage(buyerID, sellerID, sender, content)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedSeller(sellerID, chat)
	return chat, nil
}

func (s *chatServiceImpl) SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error) {
	chat := domain.NewChatMessage(buyerID, sellerID, sender, content)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedBuyer(buyerID, chat)
	return chat, nil
}

func (s *chatServiceImpl) PollMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, chatID int64) ([]*domain.ChatMessage, error) {
	latestChatTime := time.Time{}
	if chatID != 0 {
		chat, err := s.chatRepo.FindChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		if chat.BuyerID != buyerID || chat.SellerID != sellerID {
			return nil, ErrChatNotBelongToOwner
		}
		latestChatTime = chat.CreatedAt
	}

	newChats, err := s.chatRepo.FindLatestChatsByBuyerIDAndSellerID(ctx, buyerID, sellerID, latestChatTime)
	if err != nil {
		return nil, err
	}
	if len(newChats) > 0 {
		return newChats, nil
	}

	newChan := s.addSubscribedBuyers(buyerID)
	defer s.removeSubscribedBuyer(buyerID, newChan)

	select {
	case newChat := <-newChan:
		return []*domain.ChatMessage{newChat}, nil
	case <-time.After(60 * time.Second):
		return []*domain.ChatMessage{}, nil
	}
}

func (s *chatServiceImpl) PollMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, chatID int64) ([]*domain.ChatMessage, error) {
	latestChatTime := time.Time{}
	if chatID != 0 {
		chat, err := s.chatRepo.FindChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		if chat.BuyerID != buyerID || chat.SellerID != sellerID {
			return nil, ErrChatNotBelongToOwner
		}
		latestChatTime = chat.CreatedAt
	}

	newChats, err := s.chatRepo.FindLatestChatsByBuyerIDAndSellerID(ctx, buyerID, sellerID, latestChatTime)
	if err != nil {
		return nil, err
	}
	if len(newChats) > 0 {
		return newChats, nil
	}

	newChan := s.addSubscribedSellers(sellerID)
	defer s.removeSubscribedSeller(sellerID, newChan)

	select {
	case newChat := <-newChan:
		return []*domain.ChatMessage{newChat}, nil
	case <-time.After(60 * time.Second):
		return []*domain.ChatMessage{}, nil
	}
}

func (s *chatServiceImpl) notifySubscribedBuyer(buyerID int64, chat *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	for subscribedBuyer := range s.subscribedBuyers[buyerID] {
		subscribedBuyer <- chat
	}
}

func (s *chatServiceImpl) notifySubscribedSeller(sellerID int64, chat *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	for subscribedSeller := range s.subscribedSellers[sellerID] {
		subscribedSeller <- chat
	}
}

func (s *chatServiceImpl) addSubscribedBuyers(buyerID int64) chan *domain.ChatMessage {
	s.Lock()
	defer s.Unlock()

	newChan := make(chan *domain.ChatMessage, 10)

	if _, ok := s.subscribedBuyers[buyerID]; !ok {
		s.subscribedBuyers[buyerID] = make(map[chan *domain.ChatMessage]struct{})
	}
	s.subscribedBuyers[buyerID][newChan] = struct{}{}
	return newChan
}

func (s *chatServiceImpl) addSubscribedSellers(sellerID int64) chan *domain.ChatMessage {
	s.Lock()
	defer s.Unlock()

	newChan := make(chan *domain.ChatMessage, 10)

	if _, ok := s.subscribedSellers[sellerID]; !ok {
		s.subscribedSellers[sellerID] = make(map[chan *domain.ChatMessage]struct{})
	}
	s.subscribedSellers[sellerID][newChan] = struct{}{}
	return newChan
}

func (s *chatServiceImpl) removeSubscribedBuyer(buyerID int64, sub chan *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribedBuyers[buyerID], sub)
}

func (s *chatServiceImpl) removeSubscribedSeller(sellerID int64, sub chan *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribedSellers[sellerID], sub)
}
