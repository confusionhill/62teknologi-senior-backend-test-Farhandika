package dto

import "time"

type InsertBusinessDTO struct {
	Name      string  `json:"name" validate:"required"`
	Term      string  `json:"term"`
	Price     int     `json:"price"`
	Location  string  `json:"location" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Radius    float64 `json:"radius"`
}

type DeleteBusinessDTO struct {
	ID uint `json:"id" validate:"required"`
}

type InsertBusinessResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UpdateBusinessDTO struct {
	ID        uint    `json:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Term      string  `json:"term"`
	Price     int     `json:"price"`
	Location  string  `json:"location" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Radius    float64 `json:"radius"`
}

type SearchBusinessDTO struct {
	Page      uint64
	Limit     uint64
	Term      string
	Location  string
	Latitude  float64
	Longitude float64
	Locale    string
	OpenNow   bool
	OpenAT    string
}

type BusinessDTO struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Term      string     `json:"term"`
	Price     int        `json:"price"`
	Location  string     `json:"location"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Radius    float64    `json:"radius"`
	Locale    string     `json:"locale"`
	OpenNow   bool       `json:"open_now"`
	OpenAt    *time.Time `json:"open_at"`
}
