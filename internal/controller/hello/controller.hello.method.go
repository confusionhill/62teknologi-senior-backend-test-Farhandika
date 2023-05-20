package hello

import (
	"github.com/labstack/echo/v4"
)

func (c HelloController) SayHello(ctx echo.Context) (string, error) {
	return c.repo.GetHelloMessage()
}
