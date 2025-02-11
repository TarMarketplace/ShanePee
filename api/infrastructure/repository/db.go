package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shanepee.com/api/config"
	"shanepee.com/api/domain"
)

func NewDB(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&domain.User{}, &domain.ArtToy{}, &domain.PasswordChangeRequest{}); err != nil {
		return nil, err
	}
	return db, nil
}
