package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
	"shanepee.com/api/service"
)

const userIdSessionKey string = "user_id"

type AuthHandler struct {
	authSvc            service.AuthService
	defaultSessionOpts sessions.Options
}

func NewAuthHandler(authSvc service.AuthService, defaultSessionOpts sessions.Options) AuthHandler {
	return AuthHandler{
		authSvc,
		defaultSessionOpts,
	}
}

// @Summary		Register User
// @Description	Register
// @Tags			Authentication
// @Produce		json
// @Param			body	body		domain.UserCreateBody	true	"user create body"
// @Success		200		{object}	domain.User
// @Failure		400		{object}	ErrorResponse
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
	session := sessions.Default(c)
	session.Set(userIdSessionKey, data.ID)
	if err := session.Save(); err != nil {
		handleError(c, apperror.ErrInternal(err))
		return
	}
	c.JSON(http.StatusOK, data)
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary		Login User
// @Description	Login
// @Tags			Authentication
// @Produce		json
// @Param			body	body		LoginInput	true	"login input"
// @Success		200		{object}	domain.User
// @Failure		400		{object}	ErrorResponse
// @Failure		401		{object}	ErrorResponse
// @Router			/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var body LoginInput
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}
	data, err := h.authSvc.Login(c, body.Email, body.Password)
	if err != nil {
		handleError(c, err)
		return
	}
	session := sessions.Default(c)
	session.Set(userIdSessionKey, data.ID)
	if err := session.Save(); err != nil {
		handleError(c, apperror.ErrInternal(err))
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary		Logout User
// @Description	Logout
// @Tags			Authentication
// @Produce		json
// @Success		200
// @Router			/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	newSessionOpts := h.defaultSessionOpts
	newSessionOpts.MaxAge = -1
	session.Options(newSessionOpts)
	if err := session.Save(); err != nil {
		handleError(c, apperror.ErrInternal(err))
		return
	}
	c.Status(http.StatusOK)
}
