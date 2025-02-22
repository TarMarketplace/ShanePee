package domain

type CartItem struct {
	ID       int64  `json:"id" gorm:"primaryKey" example:"97"`
	OwnerID  int64  `json:"owner_id" gorm:"not null" example:"97"`
	ArtToyID int64  `json:"art_toy_id" gorm:"not null" example:"97"`
	Owner    User   `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
	ArtToy   ArtToy `gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
}

func NewCartItem(ownerID int64, artToyID int64) *CartItem {
	return &CartItem{
		ID:       GenID(),
		OwnerID:  ownerID,
		ArtToyID: artToyID,
	}
}
