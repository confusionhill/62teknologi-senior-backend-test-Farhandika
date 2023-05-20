package business

import "project-template/internal/config"

type Controller struct {
	cfg        *config.MainConfig
	repository Dependency
}

func Init(cfg *config.MainConfig, repository Dependency) *Controller {
	return &Controller{
		cfg:        cfg,
		repository: repository,
	}
}
