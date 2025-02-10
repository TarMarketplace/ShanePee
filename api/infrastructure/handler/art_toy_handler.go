package handler

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
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

//	@Summary		Create ArtToy
//	@Description	create a new ArtToy
//	@Tags			ArtToy
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.ArtToyCreateBody	true	"body of ArtToy to be created"
//	@Success		200		{object}	string
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Router			/v1/art-toy [post]
func (h *ArtToyHandler) CreateArtToy(c *gin.Context) {
	var body domain.ArtToyCreateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid request body"))
		return
	}

	artToy := domain.NewArtToy(body.Name, body.Description, body.Price, body.Photo)
	appError := h.artToySvc.CreateArtToy(c, artToy)
	if appError != nil {
		handleError(c, appError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ArtToy created successfully",
		"id":      artToy.ID,
	})
}

//	@Summary		Update ArtToy
//	@Description	Update an existing ArtToy by ID
//	@Tags			ArtToy
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64					true	"ID of the ArtToy to update"
//	@Param			body	body		domain.ArtToyUpdateBody	true	"Updated ArtToy data"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Router			/v1/art-toy/{id} [patch]
func (h *ArtToyHandler) UpdateArtToy(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid ID"))
		return
	}

	var updateBody domain.ArtToyUpdateBody
	if err := c.ShouldBindJSON(&updateBody); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid request body"))
		return
	}

	appError := h.artToySvc.UpdateArtToy(c, id, &updateBody)
	if appError != nil {
		handleError(c, appError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ArtToy updated successfully"})
}
