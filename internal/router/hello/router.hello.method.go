package hello

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-template/model/dto"
)

func (r HelloRouter) SayHello(c echo.Context) error {
	res, err := r.controller.SayHello(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "Failed",
			Data:   nil,
		})
	}
	return c.JSON(200, dto.ResponseWrapper{
		Status: 200,
		Msg:    "Success",
		Data:   res,
	})
}
