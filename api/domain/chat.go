package domain

import "time"

type ChatType string

const (
	MessageType ChatType = "MESSAGE"
	ImageType   ChatType = "IMAGE"
)

type ChatMessage struct {
	ID          int64     `json:"id" gorm:"primaryKey" example:"97"`
	SenderID    int64     `json:"sender_id" gorm:"not null" example:"97"`
	ReceiverID  int64     `json:"receiver_id" gorm:"not null" example:"97"`
	MessageType ChatType  `json:"message_type" gorm:"not null" enum:"MESSAGE,IMAGE" example:"MESSAGE"`
	Content     string    `json:"content" gorm:"not null" example:"Hello world"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type ChatList struct {
	ID                  int64     `json:"id" gorm:"primaryKey" example:"97"`
	TargetID            int64     `json:"target_id" gorm:"not null" example:"97"`
	TargetFirstName     *string   `json:"target_first_name,omitempty" example:"John"`
	TargetLastName      *string   `json:"target_last_name,omitempty" example:"Doe"`
	TargetPhoto         *string   `json:"target_photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	LastChatMessageType ChatType  `json:"last_chat_message_type" gorm:"not null" enum:"MESSAGE,IMAGE" example:"MESSAGE"`
	LastChatContent     string    `json:"last_chat_content" gorm:"not null" example:"Hello world"`
	LastChatTime        time.Time `json:"last_chat_time" gorm:"not null" example:"2021-08-01T00:00:00Z"`
}

func NewChatMessage(senderID int64, receiverID int64, messageType ChatType, content string) *ChatMessage {
	return &ChatMessage{
		ID:          GenID(),
		SenderID:    senderID,
		ReceiverID:  receiverID,
		MessageType: messageType,
		Content:     content,
		CreatedAt:   time.Now(),
	}
}

func NewChatList(ID int64, targetID int64, targetFirstName *string, targetLastName *string, targetPhoto *string, lastChatMessageType ChatType, lastChatContent string, lastChatTime time.Time) *ChatList {
	return &ChatList{
		ID:                  ID,
		TargetID:            targetID,
		TargetFirstName:     targetFirstName,
		TargetLastName:      targetLastName,
		TargetPhoto:         targetPhoto,
		LastChatMessageType: lastChatMessageType,
		LastChatContent:     lastChatContent,
		LastChatTime:        lastChatTime,
	}
}
