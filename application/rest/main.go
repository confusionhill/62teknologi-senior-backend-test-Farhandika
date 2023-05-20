package main

import (
	"errors"
	"log"
	"project-template/internal/config"
)

func main() {
	conf, err := configLoader()
	if err != nil {
		log.Fatal(err)
	}
	if err = application(conf); err != nil {
		log.Fatal(err)
	}
}

func configLoader() (*config.MainConfig, error) {
	conf, err := config.Init()
	if err != nil {
		return nil, errors.New("Cannot find secrets.config.json")
	}
	return conf, nil
}
