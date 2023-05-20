package application

import (
	"log"
	"project-template/internal/config"
	"project-template/internal/repository/auth"
	"project-template/internal/repository/business"
	"project-template/internal/repository/hello"
)

type CommonRepository struct {
	cfg          *config.MainConfig
	helloRepo    *hello.HelloRepository
	authRepo     *auth.AuthRepository
	businessRepo *business.Repository
}

func InitRepo(rsc *CommonResource, cfg *config.MainConfig) (*CommonRepository, error) {
	r := new(CommonRepository)
	r.cfg = cfg
	hRepo, err := hello.Init(rsc.Database, rsc.Redis, rsc.Mongo)
	if err != nil {
		return nil, err
	}
	r.helloRepo = hRepo
	aRepo, err := auth.Init(rsc.Database)
	if err != nil {
		return nil, err
	}
	r.authRepo = aRepo
	r.businessRepo = business.Init(rsc.Database)
	log.Println("REPOSITORY INTEGRATED!")
	return r, nil
}
