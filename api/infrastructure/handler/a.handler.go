package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
	"shanepee.com/api/service"
)

type AHandler struct {
	aSvc service.AService
}

func NewAHandler(aSvc service.AService) AHandler {
	return AHandler{
		aSvc,
	}
}

// @Summary		get A
// @Description	get A
// @Tags			a
// @Produce		json
// @Success		200	{object}	domain.A
// @Failure		400	{object}	ErrorResponse
// @Router			/v1/a [get]
func (h *AHandler) GetA(c *gin.Context) {
	data, err := h.aSvc.FindA(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		create A
// @Description	create A
// @Tags			a
// @Accept			json
// @Produce		json
// @Param			body	body		domain.ACreateBody	true	"A body"
// @Success		200		{object}	domain.A
// @Failure		400		{object}	ErrorResponse
// @Router			/v1/a [post]
func (h *AHandler) CreateA(c *gin.Context) {
	var body domain.ACreateBody
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}

	data, err := h.aSvc.CreateA(c, body)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
