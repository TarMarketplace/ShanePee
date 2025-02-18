package domain

type Cart struct {
	ID      int64   `json:"id" gorm:"primaryKey"`
	ArtToys []int64 `json:"art_toys" gorm:"type:integer[]"` // array of art toy IDs
	OwnerId int64   `json:"owner_id" gorm:"not null"`
}

func NewCart(ownerId int64) *Cart {
	return &Cart{
		ID:      GenID(),
		ArtToys: []int64{},
		OwnerId: ownerId,
	}
}
