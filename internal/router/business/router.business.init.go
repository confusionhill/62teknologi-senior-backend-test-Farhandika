package business

import (
	"github.com/labstack/echo/v4"
	"project-template/internal/config"
)

type Router struct {
	cfg        *config.MainConfig
	controller Dependency
}

func Init(cfg *config.MainConfig, controller Dependency) *Router {
	return &Router{
		cfg:        cfg,
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Echo) {
	g := e.Group("/business")
	g.POST("", r.AddBusiness)
	g.PUT("", r.EditBusiness)
	g.DELETE("", r.DeleteBusiness)
	g.GET("/search", r.GetBusinesses)
}
