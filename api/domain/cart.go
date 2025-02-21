package domain

type Cart struct {
	ID      int64 `json:"id" gorm:"primaryKey" example:"97"`
	OwnerID int64 `json:"owner_id" gorm:"not null" example:"97"`
}

type CartItem struct {
	ID       int64  `json:"id" gorm:"primaryKey" example:"97"`
	CartID   int64  `json:"cart_id" gorm:"not null" example:"97"`
	ArtToyID int64  `json:"art_toy_id" gorm:"not null" example:"97"`
	Cart     Cart   `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE;"`
	ArtToy   ArtToy `gorm:"foreignKey:ArtToyID;constraint:OnDelete:CASCADE;"`
}

func NewCart(ownerID int64) *Cart {
	return &Cart{
		ID:      GenID(),
		OwnerID: ownerID,
	}
}

func NewCartItem(cartID int64, artToyID int64) *CartItem {
	return &CartItem{
		ID:       GenID(),
		CartID:   cartID,
		ArtToyID: artToyID,
	}
}
