package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"project-template/model/dto"
)

type JwtCustomClaims struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Access uint32 `json:"access"`
	jwt.RegisteredClaims
}

type AuthGuardRequest struct {
	EchoCtx echo.Context
	Claims  *JwtCustomClaims
}

func AuthGuard(funcHandler func(a AuthGuardRequest) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.String(http.StatusBadRequest, "bad")
		}
		claims := user.Claims.(*JwtCustomClaims)
		jwt := AuthGuardRequest{
			EchoCtx: c,
			Claims:  claims,
		}
		return funcHandler(jwt)
	}
}

func TestAuthGuard(a AuthGuardRequest) error {
	return a.EchoCtx.JSON(200, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   a.Claims,
	})
}
