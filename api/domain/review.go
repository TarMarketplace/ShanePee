package domain

type Review struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Rating   int    `json:"rating" example:"5"`
	Comment  string `json:"comment" example:"This is a great art toy"`
	ArtToyID int64  `json:"art_toy_id" gorm:"not null"`
	ArtToy   ArtToy `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
}

func NewReview(rating int, comment string, artToyID int64) *Review {
	return &Review{
		ID:       GenID(),
		Rating:   rating,
		Comment:  comment,
		ArtToyID: artToyID,
	}
}
