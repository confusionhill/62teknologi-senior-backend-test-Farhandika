package config

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MainConfig struct {
	Server    ServerConfig   `json:"server"`
	Database  DatabaseConfig `json:"database"`
	Redis     RedisConfig    `json:"redis"`
	Mongo     MongoConfig    `json:"mongo"`
	JWTSecret string         `json:"jwt_secret"`
	JwtCfg    echojwt.Config
	CorsCfg   middleware.CORSConfig
	GzipConf  middleware.GzipConfig
}

type MongoConfig struct {
	Host string `json:"host"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host string `json:"host"`
}
