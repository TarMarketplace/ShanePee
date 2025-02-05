package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
	"shanepee.com/api/service"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) UserHandler {
	return UserHandler{
		userSvc,
	}
}

// @Summary		get users
// @Description	get all users
// @Tags		User
// @Produce		json
// @Success		200	{object}	[]domain.User
// @Failure		400	{object}	ErrorResponse
// @Router		/v1/user [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	data, err := h.userSvc.GetUsers(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		get user
// @Description	get user by id
// @Tags		User
// @Produce		json
// @Param		id				path			int64	true	"id of user to be fetched"
// @Success		200	{object}	domain.User
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router		/v1/user/:id [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid id"))
		return
	}

	data, err2 := h.userSvc.GetUser(c, id)
	if err2 != nil {
		handleError(c, err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// @Summary		update user
// @Description	update user by id
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		id		path		int64					true	"id of user to be updated"
// @Param		body	body		map[string]interface{}	true	"body of user to be updated"
// @Success		200		{object}	domain.User
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router		/v1/user/:id [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid id"))
		return
	}

	var body map[string]interface{}
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}

	err2 := h.userSvc.UpdateUser(c, id, body)
	if err2 != nil {
		handleError(c, err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}
