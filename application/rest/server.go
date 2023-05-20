package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"project-template/internal/config"
)

func startServer(cfg *config.MainConfig, e *echo.Echo) error {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	s := http.Server{
		Addr:    addr,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	e.HideBanner = true
	fmt.Printf("[SERVER] Server start on  %s \n", addr)
	return s.ListenAndServe()
}
