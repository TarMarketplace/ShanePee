package domain

import "time"

type Review struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Rating    int       `json:"rating" example:"5"`
	Comment   string    `json:"comment" example:"This is a great art toy"`
	OrderID   int64     `json:"order_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-08-01T00:00:00Z"`
	Order     Order     `json:"-" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

func NewReview(rating int, comment string, orderID int64) *Review {
	return &Review{
		ID:        GenID(),
		Rating:    rating,
		Comment:   comment,
		OrderID:   orderID,
		CreatedAt: time.Now(),
	}
}
