package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"shanepee.com/api/config"
)

// DefaultOptions is a wrapper for storing a gin session options with default values.
type DefaultOptions struct {
	sessions.Options
}

func NewDefaultOptions(cfg config.Config) DefaultOptions {
	defaultOpts := sessions.Options{
		Path:     "/",
		Domain:   cfg.Session.CookieDomain,
		MaxAge:   int(cfg.Session.CookieMaxAge.Seconds()),
		Secure:   cfg.Session.CookieSecure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return DefaultOptions{defaultOpts}
}
