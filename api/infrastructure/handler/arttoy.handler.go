package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
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

// @Summary		Get Art Toy by ID
// @Description	Get art toy by id
// @Tags			Art toy
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"id of art toy to be retrieved"
// @Success		200		{object}	domain.ArtToy
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router			/v1/art-toy [get]
func (h *ArtToyHandler) GetArtToyById(c *gin.Context) {
	artToyId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid art toy id"))
		return
	}

	data, appError := h.artToySvc.GetArtToyById(c, artToyId)
	if appError != nil {
		handleError(c, appError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
