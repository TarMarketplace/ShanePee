package domain

type CartItem struct {
	ID       int64  `json:"id" gorm:"primaryKey" example:"97"`
	ArtToyID int64  `json:"art_toy_id" gorm:"not null" example:"97"`
	OwnerID  int64  `json:"cart_id" gorm:"not null" example:"97"`
	ArtToy   ArtToy `json:"-" gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
	Owner    User   `json:"-" gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
}

func NewCartItem(artToyID int64, ownerID int64) *CartItem {
	return &CartItem{
		ID:       GenID(),
		ArtToyID: artToyID,
		OwnerID:  ownerID,
	}
}
