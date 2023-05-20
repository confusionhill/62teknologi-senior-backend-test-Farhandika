package business

import "gorm.io/gorm"

type Repository struct {
	Database *gorm.DB
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		Database: db,
	}
}
