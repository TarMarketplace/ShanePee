package domain

import "time"

type OrderStatus string

// TODO: Wait for the design from the frontend team
const (
	Preparing  OrderStatus = "PREPARING"
	Delivering OrderStatus = "DELIVERING"
	Completed  OrderStatus = "COMPLETED"
)

type Order struct {
	ID              int64       `json:"id" gorm:"primaryKey" example:"97"`
	TrackingNumber  *string     `json:"tracking_number,omitempty" example:"TH1234567890"`
	DeliveryService *string     `json:"delivery_service,omitempty" example:"Kerry Express"`
	SellerID        int64       `json:"seller_id" gorm:"not null" example:"97"`
	BuyerID         int64       `json:"buyer_id" gorm:"not null" example:"97"`
	Status          OrderStatus `json:"status" gorm:"not null" enum:"PREPARING,DELIVERING,COMPLETED" example:"pending"`
	CreatedAt       time.Time   `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
}

type OrderItem struct {
	ID       int64  `json:"id" gorm:"primaryKey" example:"97"`
	ArtToyID int64  `json:"art_toy_id" gorm:"not null" example:"97"`
	OrderID  int64  `json:"order_id" gorm:"not null" example:"97"`
	ArtToy   ArtToy `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
	Order    Order  `json:"-" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

func NewOrder(sellerID int64, buyerID int64) *Order {
	return &Order{
		ID:       GenID(),
		SellerID: sellerID,
		BuyerID:  buyerID,
		Status:   Preparing,
	}
}

func NewOrderItem(artToyID int64, orderID int64) *OrderItem {
	return &OrderItem{
		ID:       GenID(),
		ArtToyID: artToyID,
		OrderID:  orderID,
	}
}
