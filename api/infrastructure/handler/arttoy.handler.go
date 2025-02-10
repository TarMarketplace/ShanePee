package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/service"
)

type ArtToyHandler struct {
	artToySvc service.ArtToyService
}

func NewArtToyHandler(artToySvc service.ArtToyService) ArtToyHandler {
	return ArtToyHandler{
		artToySvc,
	}
}

// @Summary		Get Art Toys
// @Description	Get all art toys
// @Tags			Art toy
// @Accept			json
// @Produce		json
// @Success		200		{object}	domain.ArrayResponse{data=[]domain.ArtToy}
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router			/v1/art-toy [get]
func (h *ArtToyHandler) GetArtToys(c *gin.Context) {
	data, appError := h.artToySvc.GetArtToys(c)
	if appError != nil {
		handleError(c, appError)
		return
	}
	c.JSON(http.StatusOK, data)
}
