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
	GetChatList(ctx context.Context, userID int64, chatID int64, poll bool) ([]*domain.ChatList, error)
	SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error)
	SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error)
	GetChatMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, chatID int64, poll bool) ([]*domain.ChatMessage, error)
	GetChatMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, chatID int64, poll bool) ([]*domain.ChatMessage, error)
}

type chatServiceImpl struct {
	chatRepo                   domain.ChatRepository
	userRepo                   domain.UserRepository
	subscribedBuyers           map[int64]map[chan *domain.ChatMessage]struct{}
	subscribedSellers          map[int64]map[chan *domain.ChatMessage]struct{}
	subscribedUsersForChatList map[int64]map[chan *domain.ChatList]struct{}
	sync.Mutex
}

func NewChatService(chatRepo domain.ChatRepository, userRepo domain.UserRepository) ChatService {
	return &chatServiceImpl{
		chatRepo:                   chatRepo,
		userRepo:                   userRepo,
		subscribedBuyers:           make(map[int64]map[chan *domain.ChatMessage]struct{}),
		subscribedSellers:          make(map[int64]map[chan *domain.ChatMessage]struct{}),
		subscribedUsersForChatList: make(map[int64]map[chan *domain.ChatList]struct{}),
	}
}

var _ ChatService = &chatServiceImpl{}

func (s *chatServiceImpl) GetChatList(ctx context.Context, userID int64, chatID int64, poll bool) ([]*domain.ChatList, error) {
	latestChatTime := time.Time{}
	if poll {
		chat, err := s.chatRepo.FindChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		if chat.BuyerID != userID && chat.SellerID != userID {
			return nil, ErrChatNotBelongToOwner
		}
		latestChatTime = chat.CreatedAt
	}

	newChats, err := s.chatRepo.FindChatListByUserID(ctx, userID, latestChatTime)
	if err != nil {
		return nil, err
	}
	if len(newChats) > 0 {
		return newChats, nil
	}

	newChan := s.addSubscribedUserForChatList(userID)
	defer s.removeSubscribedUserForChatList(userID, newChan)

	select {
	case newChat := <-newChan:
		return []*domain.ChatList{newChat}, nil
	case <-time.After(60 * time.Second):
		return []*domain.ChatList{}, nil
	}
}

func (s *chatServiceImpl) SendMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error) {
	chat := domain.NewChatMessage(buyerID, sellerID, sender, content)

	seller, err := s.userRepo.FindUserByID(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	chatList := domain.NewChatList(chat.ID, sellerID, domain.Seller, seller.FirstName, seller.LastName, seller.Photo, content, chat.CreatedAt)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedSeller(sellerID, chat)
	s.notifySubscribedUserForChatList(sellerID, chatList)
	return chat, nil
}

func (s *chatServiceImpl) SendMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, sender domain.UserType, content string) (*domain.ChatMessage, error) {
	chat := domain.NewChatMessage(buyerID, sellerID, sender, content)

	buyer, err := s.userRepo.FindUserByID(ctx, buyerID)
	if err != nil {
		return nil, err
	}
	chatList := domain.NewChatList(chat.ID, buyerID, domain.Buyer, buyer.FirstName, buyer.LastName, buyer.Photo, content, chat.CreatedAt)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscribedBuyer(buyerID, chat)
	s.notifySubscribedUserForChatList(buyerID, chatList)
	return chat, nil
}

func (s *chatServiceImpl) GetChatMessageByBuyer(ctx context.Context, buyerID int64, sellerID int64, chatID int64, poll bool) ([]*domain.ChatMessage, error) {
	latestChatTime := time.Time{}
	if poll {
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

func (s *chatServiceImpl) GetChatMessageBySeller(ctx context.Context, buyerID int64, sellerID int64, chatID int64, poll bool) ([]*domain.ChatMessage, error) {
	latestChatTime := time.Time{}
	if poll {
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
	delete(s.subscribedSellers, buyerID)
}

func (s *chatServiceImpl) notifySubscribedSeller(sellerID int64, chat *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	for subscribedSeller := range s.subscribedSellers[sellerID] {
		subscribedSeller <- chat
	}
	delete(s.subscribedSellers, sellerID)
}

func (s *chatServiceImpl) notifySubscribedUserForChatList(userID int64, chat *domain.ChatList) {
	s.Lock()
	defer s.Unlock()

	for subscribedUser := range s.subscribedUsersForChatList[userID] {
		subscribedUser <- chat
	}
	delete(s.subscribedUsersForChatList, userID)
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

func (s *chatServiceImpl) addSubscribedUserForChatList(userID int64) chan *domain.ChatList {
	s.Lock()
	defer s.Unlock()

	newChan := make(chan *domain.ChatList, 10)

	if _, ok := s.subscribedUsersForChatList[userID]; !ok {
		s.subscribedUsersForChatList[userID] = make(map[chan *domain.ChatList]struct{})
	}
	s.subscribedUsersForChatList[userID][newChan] = struct{}{}
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

func (s *chatServiceImpl) removeSubscribedUserForChatList(userID int64, sub chan *domain.ChatList) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribedUsersForChatList[userID], sub)
}
