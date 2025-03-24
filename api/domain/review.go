package domain

import "time"

type Review struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Rating    int       `json:"rating" example:"5"`
	Comment   string    `json:"comment" example:"This is a great art toy"`
	ArtToyID  int64     `json:"art_toy_id" gorm:"not null"`
	ArtToy    ArtToy    `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
	UserID    int64     `json:"user_id" gorm:"not null"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReview(rating int, comment string, artToyID int64, userID int64) *Review {
	return &Review{
		ID:       GenID(),
		Rating:   rating,
		Comment:  comment,
		ArtToyID: artToyID,
		UserID:   userID,
	}
}

type ReviewResponse struct {
	ID        int64  `json:"id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	FirstName string `json:"first_name"`
	Photo     string `json:"photo"`
	Date      string `json:"date"`
}

func NewReviewResponse(review *Review) *ReviewResponse {
	firstName := ""
	if review.User.FirstName != nil {
		firstName = *review.User.FirstName
		if len(firstName) > 3 {
			firstName = firstName[:3] + "****"
		}
	}

	photo := ""
	if review.User.Photo != nil {
		photo = *review.User.Photo
	}

	return &ReviewResponse{
		ID:        review.ID,
		Rating:    review.Rating,
		Comment:   review.Comment,
		FirstName: firstName,
		Photo:     photo,
		Date:      review.CreatedAt.Format(time.RFC3339),
	}
}
