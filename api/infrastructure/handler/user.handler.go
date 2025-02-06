package handler

import (
	"net/http"

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

// @Summary		Update User
// @Description	update user by id
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		body	body		map[string]interface{}	true	"body of user to be updated"
// @Success		200		{object}	domain.User
// @Failure		400		{object}	ErrorResponse
// @Failure		404		{object}	ErrorResponse
// @Router		/v1/user [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var id int64
	// TODO: get id from cookie

	var body map[string]interface{}
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}

	appError := h.userSvc.UpdateUser(c, id, body)
	if appError != nil {
		handleError(c, appError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
