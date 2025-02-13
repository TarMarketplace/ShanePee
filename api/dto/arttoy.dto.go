package dto

type ArtToyCreateBody struct {
	Name        string  `json:"name" example:"Tuna"`
	Description string  `json:"description" example:"Delicious fish"`
	Price       float64 `json:"price" example:"55.55"`
	Photo       *string `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..." nullable:"true"`
}

type ArtToyUpdateBody struct {
	Name         *string  `json:"name,omitempty" example:"Tuna"`
	Description  *string  `json:"description,omitempty" example:"Delicious fish"`
	Price        *float64 `json:"price,omitempty" example:"55.55"`
	Photo        *string  `json:"photo,omitempty" example:"data:image/png;base64,mfkirjIDSFIj324if..."`
	Availability *bool    `json:"availability,omitempty" example:"false"`
}

func (b *ArtToyUpdateBody) ToMap() map[string]any {
	result := make(map[string]any)

	if b.Name != nil {
		result["name"] = b.Name
	}
	if b.Description != nil {
		result["description"] = b.Description
	}
	if b.Price != nil {
		result["price"] = b.Price
	}
	if b.Photo != nil {
		result["photo"] = b.Photo
	}
	if b.Availability != nil {
		result["availability"] = b.Availability
	}

	return result
}
