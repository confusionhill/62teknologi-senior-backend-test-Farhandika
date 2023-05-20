package hello

import (
	hc "project-template/internal/controller/hello"
)

type HelloRouter struct {
	controller *hc.HelloController
}

func Init(c *hc.HelloController) *HelloRouter {
	r := &HelloRouter{controller: c}
	return r
}
