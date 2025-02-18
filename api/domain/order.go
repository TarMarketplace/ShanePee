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
	SellerID        int64       `json:"seller_id" gorm:"not null" example:"97"`
	BuyerID         int64       `json:"buyer_id" gorm:"not null" example:"97"`
	Status          OrderStatus `json:"status" gorm:"not null" enum:"PENDING,SHIPPING,COMPLETED" example:"pending"`
	CreatedAt       time.Time   `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type OrderItem struct {
	ID       int64 `json:"id" gorm:"primaryKey" example:"97"`
	ArtToyID int64 `json:"art_toy_id" gorm:"not null" example:"97"`
	OrderID  int64 `json:"order_id" gorm:"not null" example:"97"`
}

func NewOrder(artToys []int64, sellerID int64, buyerID int64) *Order {
	return &Order{
		ID:       GenID(),
		SellerID: sellerID,
		BuyerID:  buyerID,
		Status:   Pending,
	}
}

func NewOrderItem(artToyID int64, orderID int64) *OrderItem {
	return &OrderItem{
		ID:       GenID(),
		ArtToyID: artToyID,
		OrderID:  orderID,
	}
}
