package domain

import "time"

type OrderStatus string

// TODO: Wait for the design from the frontend team
const (
	Pending   OrderStatus = "pending"
	Shipping  OrderStatus = "shipping"
	Completed OrderStatus = "completed"
)

type Order struct {
	ID              int64       `json:"id" gorm:"primaryKey" example:"97"`
	ArtToys         []int64     `json:"art_toys" gorm:"type:integer[]"` // array of art toy IDs
	TrackingNumber  *string     `json:"tracking_number" example:"TH1234567890"`
	DeliveryService *string     `json:"delivery_service" example:"Kerry Express"`
	SellerId        int64       `json:"seller_id" gorm:"not null" example:"97"`
	BuyerId         int64       `json:"buyer_id" gorm:"not null" example:"97"`
	Status          OrderStatus `json:"status" gorm:"not null" example:"pending"`
	CreatedAt       time.Time   `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
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
