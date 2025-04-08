package domain

import "time"

type ArtToy struct {
	ID            int64       `json:"id" gorm:"primaryKey"`
	Name          string      `json:"name" gorm:"not null"`
	Description   string      `json:"description" gorm:"not null"`
	Price         float64     `json:"price" gorm:"not null"`
	Photo         *string     `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	Availability  bool        `json:"availability" gorm:"not null"`
	OwnerID       int64       `json:"owner_id" gorm:"not null"`
	ReleaseDate   time.Time   `json:"release_date" gorm:"not null" example:"2021-01-01T00:00:00Z"`
	AverageRating float64     `json:"average_rating" gorm:"-:migration;->"`
	Owner         User        `json:"owner" gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
	OrderItems    []OrderItem `json:"-" gorm:"foreignKey:ArtToyID"`
}

type ArtToySortKey int

const (
	ArtToyPriceSortKey ArtToySortKey = iota
	ArtToyReleaseDateSortKey
)

type ArtToySearchParams struct {
	Keyword string
	SortKey *ArtToySortKey
	Reverse bool
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
