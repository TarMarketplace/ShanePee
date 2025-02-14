package arttoy

import (
	"shanepee.com/api/service"
)

type ArtToyHandler struct {
	artToySvc service.ArtToyService
}

func NewHandler(artToySvc service.ArtToyService) ArtToyHandler {
	return ArtToyHandler{
		artToySvc,
	}
}
