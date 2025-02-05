package handler

import (
	"net/http"
	"strconv"

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
// @Tags		A
// @Produce		json
// @Success		200	{object}	domain.A
// @Failure		400	{object}	ErrorResponse
// @Router		/v1/a [get]
func (h *AHandler) GetA(c *gin.Context) {
	data, err := h.aSvc.FindManyA(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		get A
// @Description	get A
// @Tags		A
// @Produce		json
// @Param		id		path		int64	true	"ID"
// @Success		200		{object}	domain.A
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router		/v1/a/:id [get]
func (h *AHandler) GetAById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid ID"))
		return
	}

	data, err2 := h.aSvc.FindOneA(c, id)
	if err2 != nil {
		handleError(c, err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		create A
// @Description	create A
// @Tags		A
// @Accept		json
// @Produce		json
// @Param		body	body		domain.ACreateBody	true	"A body"
// @Success		200		{object}	domain.A
// @Failure		400		{object}	ErrorResponse
// @Router		/v1/a [post]
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

// @Summary		update A
// @Description	update A
// @Tags		A
// @Accept		json
// @Produce		json
// @Param		id		path		int64					true	"ID"
// @Param		body	body		map[string]interface{}	true	"A body"
// @Success		200		{object}	domain.A
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router		/v1/a/:id [patch]
func (h *AHandler) UpdateA(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid ID"))
		return
	}

	var body map[string]interface{}
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}

	data, err2 := h.aSvc.UpdateA(c, id, body)
	if err2 != nil {
		handleError(c, err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		delete A
// @Description	delete A
// @Tags		A
// @Accept		json
// @Produce		json
// @Param		id		path		int64	true	"ID"
// @Success		200		{object}	string
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router		/v1/a/:id [delete]
func (h *AHandler) DeleteA(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid ID"))
		return
	}

	err2 := h.aSvc.DeleteA(c, id)
	if err2 != nil {
		handleError(c, err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}
