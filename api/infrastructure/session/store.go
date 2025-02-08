package session

import (
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"gorm.io/gorm"
	"shanepee.com/api/config"
)

func NewStore(cfg config.Config, defaultOptions sessions.Options, db *gorm.DB) sessions.Store {
	store := gormsessions.NewStore(db, true, []byte(cfg.SessionConfig.Key))
	store.Options(defaultOptions)
	return store
}
