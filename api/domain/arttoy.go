package domain

type ArtToy struct {
	ID           int64   `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"not null"`
	Description  string  `json:"description" gorm:"not null"`
	Price        float64 `json:"price" gorm:"not null"`
	Photo        *string `json:"photo"`
	Availability bool    `json:"availability" gorm:"not null"`
	OwnerId      int64   `json:"owner_id" gorm:"not null"`
	// TODO: add more fields about review, rating, etc.
}

type ArtToyCreateBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       *string `json:"photo"`
}

type ArtToyUpdateBody struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Photo        *string `json:"photo"`
	Availability bool    `json:"availability"`
}

func NewArtToy(name string, description string, price float64, photo *string, ownerId int64) *ArtToy {
	return &ArtToy{
		ID:           GenID(),
		Name:         name,
		Description:  description,
		Price:        price,
		Photo:        photo,
		Availability: true,
		OwnerId:      ownerId,
	}
}
