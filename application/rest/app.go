package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	app "project-template/application"
	"project-template/internal/config"
	"project-template/internal/router/health"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func application(cfg *config.MainConfig) error {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	rsc, err := app.InitResource(cfg)
	if err != nil {
		return err
	}
	r, err := app.InitRepo(rsc, cfg)
	if err != nil {
		return err
	}
	c, err := app.InitController(r, cfg)
	if err != nil {
		return err
	}
	router, err := app.InitRouter(c, cfg)
	if err != nil {
		return err
	}

	e.GET("/health", health.CheckHealth)
	e.Use(middleware.CORSWithConfig(cfg.CorsCfg))
	e.Use(middleware.GzipWithConfig(cfg.GzipConf))

	// register business router
	router.BusinessRouter.Register(e)
	return startServer(cfg, e)
}
