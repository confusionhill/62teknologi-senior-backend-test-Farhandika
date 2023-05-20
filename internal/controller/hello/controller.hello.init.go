package hello

import hr "project-template/internal/repository/hello"

type HelloController struct {
	repo HelloRepositoryInterface
}

func Init(r *hr.HelloRepository) *HelloController {
	c := &HelloController{repo: r}
	return c
}
