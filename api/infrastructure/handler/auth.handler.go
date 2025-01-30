package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
	"shanepee.com/api/service"
)

type AuthHandler struct {
  authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) AuthHandler {
	return AuthHandler{
		authSvc,
	}
}

// @Summary		Register User
// @Description	Register
// @Tags			Authentication
// @Produce		json
// @Param body body domain.UserCreateBody true "user create body"
// @Success		200	{object}	domain.User
// @Failure		400	{object}	ErrorResponse
// @Router			/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
  var body domain.UserCreateBody
  // TODO: validate body
  if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
  }
	data, err := h.authSvc.Register(c, body.Email, body.Password)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
