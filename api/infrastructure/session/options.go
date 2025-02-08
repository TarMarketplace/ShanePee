package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"shanepee.com/api/config"
)

// NewOptions is a function for creating a gin session options with default values from config.
func NewOptions(cfg config.Config) sessions.Options {
	return sessions.Options{
		Path:     "/",
		Domain:   cfg.Session.CookieDomain,
		MaxAge:   int(cfg.Session.CookieMaxAge.Seconds()),
		Secure:   cfg.Session.CookieSecure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
