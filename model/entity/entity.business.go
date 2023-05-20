package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Term      string
	Price     int
	Location  string
	Latitude  float64
	Longitude float64
	Radius    float64
	Locale    string
	OpenNow   bool
	OpenAt    *sql.NullTime
}

type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null"`
}
