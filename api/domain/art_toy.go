package domain

type ArtToy struct {
	ID     int64   `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"not null"`
	Detail string  `json:"detail" gorm:"not null"`
	Price  float64 `json:"price" gorm:"not null"`
	// TODO: add more fields about review, rating, etc.
}
