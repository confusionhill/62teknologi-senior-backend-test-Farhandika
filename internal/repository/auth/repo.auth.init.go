package auth

import (
	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func Init(database *gorm.DB) (*AuthRepository, error) {
	r := &AuthRepository{
		Database: database,
	}
	return r, nil
}
