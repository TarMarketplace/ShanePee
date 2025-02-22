package domain

import "time"

type ArtToy struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	Price        float64   `json:"price" gorm:"not null"`
	Photo        *string   `json:"photo" nullable:"true" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	Availability bool      `json:"availability" gorm:"not null"`
	OwnerID      int64     `json:"owner_id" gorm:"not null"`
	ReleaseDate  time.Time `json:"release_date" gorm:"not null" example:"2021-01-01T00:00:00Z"`
}

type Review struct {
	ID       int64   `json:"id" gorm:"primaryKey"`
	Rating   *int    `json:"rating" example:"5"`
	Comment  *string `json:"comment" example:"Good toy"`
	ArtToyID int64   `json:"art_toy_id" gorm:"not null"`
	ArtToy   ArtToy  `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
}

func NewArtToy(name string, description string, price float64, photo *string, ownerID int64) *ArtToy {
	return &ArtToy{
		ID:           GenID(),
		Name:         name,
		Description:  description,
		Price:        price,
		Photo:        photo,
		Availability: true,
		OwnerID:      ownerID,
		ReleaseDate:  time.Now(),
	}
}
