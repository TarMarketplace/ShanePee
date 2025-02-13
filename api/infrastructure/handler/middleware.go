package handler

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	ContextSessionKey = "__session"
)

// Huma abstract over api framework so, it not easy to get
// gin session from handler. The workaround is to use gin middleware to
// extract session before it would be abstract into huma.
func GetUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		newCtx := context.WithValue(ctx.Request.Context(), ContextSessionKey, session)

		ctx.Request = ctx.Request.WithContext(newCtx)

		ctx.Next()
	}
}

func GetSession(ctx context.Context) sessions.Session {
	v := ctx.Value(ContextSessionKey)
	if v == nil {
		return nil
	}
	if session, ok := v.(sessions.Session); ok {
		return session
	}
	return nil
}

// Return nil when user does not authenticated
func GetUserID(ctx context.Context) *int64 {
	session := GetSession(ctx)
	if session == nil {
		return nil
	}
	id := session.Get(userIdSessionKey)
	if id == nil {
		return nil
	}
	userId := id.(int64)
	return &userId
}
