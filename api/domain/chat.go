package domain

import "time"

type ChatMessage struct {
	ID         int64     `json:"id" gorm:"primaryKey" example:"97"`
	SenderID   int64     `json:"sender_id" gorm:"not null" example:"97"`
	ReceiverID int64     `json:"receiver_id" gorm:"not null" example:"97"`
	Content    string    `json:"content" gorm:"not null" example:"Hello world"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type ChatList struct {
	ID              int64     `json:"id" gorm:"primaryKey" example:"97"`
	TargetID        int64     `json:"target_id" gorm:"not null" example:"97"`
	TargetFirstName *string   `json:"target_first_name,omitempty" example:"John"`
	TargetLastName  *string   `json:"target_last_name,omitempty" example:"Doe"`
	TargetPhoto     *string   `json:"target_photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	LastChatMessage string    `json:"last_chat_message" gorm:"not null" example:"Hello world"`
	LastChatTime    time.Time `json:"last_chat_time" gorm:"not null" example:"2021-08-01T00:00:00Z"`
}

func NewChatMessage(senderID int64, receiverID int64, content string) *ChatMessage {
	return &ChatMessage{
		ID:         GenID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
	}
}

func NewChatList(ID int64, targetID int64, targetFirstName *string, targetLastName *string, targetPhoto *string, lastChatMessage string, lastChatTime time.Time) *ChatList {
	return &ChatList{
		ID:              ID,
		TargetID:        targetID,
		TargetFirstName: targetFirstName,
		TargetLastName:  targetLastName,
		TargetPhoto:     targetPhoto,
		LastChatMessage: lastChatMessage,
		LastChatTime:    lastChatTime,
	}
}
