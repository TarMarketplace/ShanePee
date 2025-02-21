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
	db.Exec("PRAGMA foreign_keys = ON;")
	if err = db.AutoMigrate(&domain.User{}, &domain.ArtToy{}, &domain.Cart{}, &domain.CartItem{}, &domain.Order{}, &domain.OrderItem{}, &domain.PasswordResetRequest{}); err != nil {
		return nil, err
	}
	return db, nil
}
