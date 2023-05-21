package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"project-template/internal/config"
	"syscall"
	"time"
)

func startServer(cfg *config.MainConfig, e *echo.Echo) error {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	server := http.Server{
		Addr:    addr,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	e.HideBanner = true

	ctx := context.Background()
	signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	go func() {
		fmt.Printf("[SERVER] Server start on  %s \n", addr)
		server.ListenAndServe()
	}()
	<-ctx.Done()
	time.Sleep(10 * time.Second)
	fmt.Println("Server shutting down...")
	return server.Shutdown(context.Background())
}
