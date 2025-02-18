package domain

type Cart struct {
	ID      int64 `json:"id" gorm:"primaryKey" example:"97"`
	OwnerId int64 `json:"owner_id" gorm:"not null" example:"97"`
}

type CartItem struct {
	ID       int64 `json:"id" gorm:"primaryKey" example:"97"`
	CartID   int64 `json:"cart_id" gorm:"not null" example:"97"`
	ArtToyId int64 `json:"art_toy_id" gorm:"not null" example:"97"`
}

func NewCart(ownerId int64) *Cart {
	return &Cart{
		ID:      GenID(),
		OwnerId: ownerId,
	}
}

func NewCartItem(cartId int64, artToyId int64) *CartItem {
	return &CartItem{
		ID:       GenID(),
		CartID:   cartId,
		ArtToyId: artToyId,
	}
}
