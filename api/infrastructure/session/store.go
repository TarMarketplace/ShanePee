package session

import (
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"gorm.io/gorm"
	"shanepee.com/api/config"
)

func NewStore(cfg config.Config, defaultOptions DefaultOptions, db *gorm.DB) sessions.Store {
	store := gormsessions.NewStore(db, true, []byte(cfg.Session.Key))
	store.Options(defaultOptions.Options)
	return store
}
