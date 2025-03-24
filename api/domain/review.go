package domain

import "time"

type Review struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Rating    int       `json:"rating" example:"5"`
	Comment   string    `json:"comment" example:"This is a great art toy"`
	CreatedAt time.Time `json:"created_at" gorm:"not null" example:"2021-01-01T00:00:00Z"`
	ArtToyID  int64     `json:"art_toy_id" gorm:"not null"`
	ArtToy    ArtToy    `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
}

type ReviewWithTruncatedBuyer struct {
	Rating         int     `json:"rating" example:"5"`
	Comment        string  `json:"comment" example:"This is a great art toy"`
	BuyerFirstName *string `json:"buyer_first_name,omitempty" example:"John"`
	BuyerLastName  *string `json:"buyer_last_name,omitempty" example:"Doe"`
	BuyerPhoto     *string `json:"buyer_photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
}

func NewReview(rating int, comment string, artToyID int64) *Review {
	return &Review{
		ID:        GenID(),
		Rating:    rating,
		Comment:   comment,
		ArtToyID:  artToyID,
		CreatedAt: time.Now(),
	}
}
