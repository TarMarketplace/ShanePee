package domain

import "time"

type OrderStatus string

// TODO: Wait for the design from the frontend team
const (
	Pending   OrderStatus = "PENDING"
	Shipping  OrderStatus = "SHIPPING"
	Completed OrderStatus = "COMPLETED"
)

type Order struct {
	ID              int64       `json:"id" gorm:"primaryKey" example:"97"`
	TrackingNumber  *string     `json:"tracking_number" example:"TH1234567890"`
	DeliveryService *string     `json:"delivery_service" example:"Kerry Express"`
	SellerId        int64       `json:"seller_id" gorm:"not null" example:"97"`
	BuyerId         int64       `json:"buyer_id" gorm:"not null" example:"97"`
	Status          OrderStatus `json:"status" gorm:"not null" enum:"PENDING,SHIPPING,COMPLETED" example:"pending"`
	CreatedAt       time.Time   `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type OrderItem struct {
	ID       int64 `json:"id" gorm:"primaryKey" example:"97"`
	ArtToyId int64 `json:"art_toy_id" gorm:"not null" example:"97"`
	OrderId  int64 `json:"order_id" gorm:"not null" example:"97"`
}

func NewOrder(sellerId int64, buyerId int64) *Order {
	return &Order{
		ID:       GenID(),
		SellerId: sellerId,
		BuyerId:  buyerId,
		Status:   Pending,
	}
}

func NewOrderItem(artToyId int64, orderId int64) *OrderItem {
	return &OrderItem{
		ID:       GenID(),
		ArtToyId: artToyId,
		OrderId:  orderId,
	}
}
