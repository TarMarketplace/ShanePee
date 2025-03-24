package domain

import "time"

type Chat struct {
	ID        int64     `json:"id" gorm:"primaryKey" example:"97"`
	BuyerID   int64     `json:"buyer_id" gorm:"not null" example:"97"`
	SellerID  int64     `json:"seller_id" gorm:"not null" example:"97"`
	Sender    UserType  `json:"sender" gorm:"not null" enum:"BUYER,SELLER" example:"BUYER"`
	Meassage  string    `json:"message" gorm:"not null" example:"Hello world"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type ChatList struct {
	ID               int64     `json:"id" gorm:"primaryKey" example:"97"`
	TargetID         int64     `json:"buyer_id" gorm:"not null" example:"97"`
	TargetType       int64     `json:"target_type" gorm:"not null" enum:"BUYER,SELLER" example:"BUYER"`
	TargetFirstName  *string   `json:"target_first_name,omitempty" example:"John"`
	TargetLastName   *string   `json:"target_last_name,omitempty" example:"Doe"`
	TargetPhoto      *string   `json:"target_photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	LastChatMeassage string    `json:"last_chat_message" gorm:"not null" example:"Hello world"`
	LastChatTime     time.Time `json:"last_chat_time" gorm:"not null" example:"2021-08-01T00:00:00Z"`
}

func NewChat(buyerID int64, sellerID int64, sender UserType, message string) *Chat {
	return &Chat{
		ID:        GenID(),
		BuyerID:   buyerID,
		SellerID:  sellerID,
		Sender:    sender,
		Meassage:  message,
		CreatedAt: time.Now(),
	}
}
