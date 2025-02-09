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

type RequestPasswordChangeInput struct {
	Email string `json:"email"`
}

// @Summary		Request a password reset
// @Description	Initiates a password reset process by sending an email with reset instructions
// @Tags			Authentication
// @Param			body	body	RequestPasswordChangeInput	true	"input"
// @Success		200
// @Router			/v1/auth/password-change-requests [post]
func (h *AuthHandler) CreatePasswordChangeRequests(c *gin.Context) {
	var body RequestPasswordChangeInput
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}
	if err := h.authSvc.RequestPasswordChange(c, body.Email); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

type ChangePasswordInput struct {
	RequestID   int64  `json:"request_id"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

// @Summary		Change password
// @Description	Change password of a user using token and request id
// @Tags			Authentication
// @Param			body	body	ChangePasswordInput	true	"input"
// @Success		200
// @Failure		401	{object}	ErrorResponse
// @Router			/v1/auth/change-password [post]
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var body ChangePasswordInput
	if err := c.ShouldBind(&body); err != nil {
		handleError(c, apperror.ErrBadRequest("Invalid body"))
		return
	}
	if err := h.authSvc.ChangePassword(c, body.RequestID, body.Token, body.NewPassword); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

// @Summary		Get current authenticated user
// @Description	Get authenticated user from the session
// @Tags			Authentication
// @Produce		json
// @Success		200	{object}	domain.User
// @Failure		401	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Router			/v1/auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	var userId int64
	session := sessions.Default(c)
	id := session.Get(userIdSessionKey)
	if id == nil {
		handleError(c, apperror.ErrUnauthorized("Authentication required"))
		return
	}
	userId = id.(int64)

	data, err := h.authSvc.GetUserByID(c, userId)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
