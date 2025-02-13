package dto

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

func (b *ArtToyUpdateBody) ToMap() map[string]any {
	panic("Unimplemented")
}
