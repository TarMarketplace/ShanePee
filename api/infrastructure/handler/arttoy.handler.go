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

// @Summary      Create Art toy
// @Description  Create a new art toy record
// @Tags         Art toy
// @Produce      json
// @Param        body  body  domain.ArtToyCreateBody  true  "body of Art toy to be created"
// @Success      200   {object}  domain.ArtToy
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /v1/arttoy [post]
func (h *ArtToyHandler) CreateArtToy(c *gin.Context) {
	var body domain.ArtToyCreateBody
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}

	artToy := domain.NewArtToy(body.Name, body.Description, body.Price, body.Photo, body.OwnerId)
	err := h.artToySvc.CreateArtToy(c, artToy)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, artToy)
}

// @Summary		Update Art toy
// @Description	Update an existing art toy by ID
// @Tags			Art toy
// @Accept			json
// @Produce		json
// @Param			id		path		int64					true	"ID of the art toy to update"
// @Param			body	body		domain.ArtToyUpdateBody	true	"Updated art toy data"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router			/v1/arttoy/{id} [put]
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

	c.JSON(http.StatusOK, updateBody)
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
