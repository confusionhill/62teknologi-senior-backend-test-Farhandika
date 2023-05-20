package config

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"os"
	locMid "project-template/internal/middleware"
)

func Init() (*MainConfig, error) {
	jsonFile, err := os.Open("files/secrets/secrets.config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(MainConfig)
	json.Unmarshal(byteValue, conf)
	conf.JwtCfg = echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(locMid.JwtCustomClaims)
		},
		SigningKey: []byte(conf.JWTSecret),
	}
	conf.CorsCfg = middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		MaxAge:       0,
	}
	conf.GzipConf = middleware.GzipConfig{
		Level: 5,
	}
	return conf, nil
}
