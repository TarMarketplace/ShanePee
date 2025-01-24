package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.A{})
	return db, nil
}
