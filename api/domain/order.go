package domain

import "time"

type Order struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	ArtToys         []int64   `json:"art_toys" gorm:"type:integer[]"` // array of art toy IDs
	TrackingNumber  *string   `json:"tracking_number" example:"TH1234567890"`
	DeliveryService *string   `json:"delivery_service"`
	SellerId        int64     `json:"seller_id" gorm:"not null"`
	BuyerId         int64     `json:"buyer_id" gorm:"not null"`
	Status          string    `json:"status" gorm:"not null" example:"pending"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewOrder(artToys []int64, sellerId int64, buyerId int64) *Order {
	return &Order{
		ID:       GenID(),
		ArtToys:  artToys,
		SellerId: sellerId,
		BuyerId:  buyerId,
		Status:   "pending",
	}
}
