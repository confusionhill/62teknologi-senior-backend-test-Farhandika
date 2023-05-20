package auth

import (
	"project-template/internal/config"
	ac "project-template/internal/controller/auth"
)

type AuthRouter struct {
	controller *ac.AuthController
	cfg        *config.MainConfig
}

func Init(c *ac.AuthController, cfg *config.MainConfig) *AuthRouter {
	r := &AuthRouter{controller: c, cfg: cfg}
	return r
}
