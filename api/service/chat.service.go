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
	GetChatList(ctx context.Context, userID int64, poll bool, chatID int64) ([]*domain.ChatList, error)
	SendMessage(ctx context.Context, senderID int64, receiverID int64, content string) (*domain.ChatMessage, error)
	GetChatMessage(ctx context.Context, senderID int64, receiverID int64, poll bool, chatID int64) ([]*domain.ChatMessage, error)
}

type chatServiceImpl struct {
	chatRepo               domain.ChatRepository
	userRepo               domain.UserRepository
	subscribers            map[int64]map[chan *domain.ChatMessage]struct{}
	subscribersForChatList map[int64]map[chan *domain.ChatList]struct{}
	sync.Mutex
}

func NewChatService(chatRepo domain.ChatRepository, userRepo domain.UserRepository) ChatService {
	return &chatServiceImpl{
		chatRepo:               chatRepo,
		userRepo:               userRepo,
		subscribers:            make(map[int64]map[chan *domain.ChatMessage]struct{}),
		subscribersForChatList: make(map[int64]map[chan *domain.ChatList]struct{}),
	}
}

var _ ChatService = &chatServiceImpl{}

func (s *chatServiceImpl) GetChatList(ctx context.Context, userID int64, poll bool, chatID int64) ([]*domain.ChatList, error) {
	latestChatTime := time.Time{}
	if poll {
		chat, err := s.chatRepo.FindChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		if chat.SenderID != userID && chat.ReceiverID != userID {
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

	newChan := s.addSubscriberForChatList(userID)
	defer s.removeSubscriberForChatList(userID, newChan)

	select {
	case newChat := <-newChan:
		return []*domain.ChatList{newChat}, nil
	case <-time.After(60 * time.Second):
		return []*domain.ChatList{}, nil
	}
}

func (s *chatServiceImpl) SendMessage(ctx context.Context, senderID int64, receiverID int64, content string) (*domain.ChatMessage, error) {
	chat := domain.NewChatMessage(senderID, receiverID, content)

	receiver, err := s.userRepo.FindUserByID(ctx, receiverID)
	if err != nil {
		return nil, err
	}
	chatList := domain.NewChatList(chat.ID, receiverID, receiver.FirstName, receiver.LastName, receiver.Photo, content, chat.CreatedAt)
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil, err
	}

	s.notifySubscriber(senderID, chat)
	s.notifySubscriber(receiverID, chat)
	s.notifySubscriberForChatList(receiverID, chatList)
	return chat, nil
}

func (s *chatServiceImpl) GetChatMessage(ctx context.Context, receiverID int64, senderID int64, poll bool, chatID int64) ([]*domain.ChatMessage, error) {
	latestChatTime := time.Time{}
	if poll {
		chat, err := s.chatRepo.FindChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		if !((chat.ReceiverID == receiverID && chat.SenderID == senderID) || (chat.ReceiverID == senderID && chat.SenderID == receiverID)) {
			return nil, ErrChatNotBelongToOwner
		}
		latestChatTime = chat.CreatedAt
	}

	newChats, err := s.chatRepo.FindLatestChatsBySenderIDAndReceiverID(ctx, senderID, receiverID, latestChatTime)
	if err != nil {
		return nil, err
	}
	if len(newChats) > 0 {
		return newChats, nil
	}

	newChan := s.addSubscriber(receiverID)
	defer s.removeSubscriber(receiverID, newChan)

	select {
	case newChat := <-newChan:
		return []*domain.ChatMessage{newChat}, nil
	case <-time.After(60 * time.Second):
		return []*domain.ChatMessage{}, nil
	}
}

func (s *chatServiceImpl) notifySubscriber(userID int64, chat *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	for subscriber := range s.subscribers[userID] {
		subscriber <- chat
	}
	delete(s.subscribers, userID)
}

func (s *chatServiceImpl) notifySubscriberForChatList(userID int64, chat *domain.ChatList) {
	s.Lock()
	defer s.Unlock()

	for subscriber := range s.subscribersForChatList[userID] {
		subscriber <- chat
	}
	delete(s.subscribersForChatList, userID)
}

func (s *chatServiceImpl) addSubscriber(userID int64) chan *domain.ChatMessage {
	s.Lock()
	defer s.Unlock()

	newChan := make(chan *domain.ChatMessage, 10)

	if _, ok := s.subscribers[userID]; !ok {
		s.subscribers[userID] = make(map[chan *domain.ChatMessage]struct{})
	}
	s.subscribers[userID][newChan] = struct{}{}
	return newChan
}

func (s *chatServiceImpl) addSubscriberForChatList(userID int64) chan *domain.ChatList {
	s.Lock()
	defer s.Unlock()

	newChan := make(chan *domain.ChatList, 10)

	if _, ok := s.subscribersForChatList[userID]; !ok {
		s.subscribersForChatList[userID] = make(map[chan *domain.ChatList]struct{})
	}
	s.subscribersForChatList[userID][newChan] = struct{}{}
	return newChan
}

func (s *chatServiceImpl) removeSubscriber(userID int64, sub chan *domain.ChatMessage) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribers[userID], sub)
}

func (s *chatServiceImpl) removeSubscriberForChatList(userID int64, sub chan *domain.ChatList) {
	s.Lock()
	defer s.Unlock()

	delete(s.subscribersForChatList[userID], sub)
}
