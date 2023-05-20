package auth

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"project-template/internal/middleware"
	"project-template/model/dto"
	"project-template/model/entity"
	"strings"
)

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type LoginUserDTOResponse struct {
	AccessToken string      `json:"access_token"`
	Info        interface{} `json:"user_info"`
}

type RegisterUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (r *AuthRouter) LoginUser(c echo.Context) error {
	u := new(LoginUserDTO)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(u); err != nil {
		msg := "Missing Field(s)"
		msgErr := err.Error()
		if strings.Contains("Email", msgErr) {
			msg = "Email not recognized"
		}
		return c.String(http.StatusBadRequest, msg)
	}
	res, err := r.controller.LoginUser(&entity.User{
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ResponseWrapper{
			Status: http.StatusNotFound,
			Msg:    "Account not found!",
			Data:   nil,
		})
	}
	claims := &middleware.JwtCustomClaims{
		Id:     res.Id,
		Name:   res.FullName,
		Access: res.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: nil,
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(r.cfg.JWTSecret))
	if err != nil {
		log.Error(err)
	}
	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "Ok",
		Data: LoginUserDTOResponse{
			AccessToken: t,
			Info: dto.UserInformationDTO{
				Id:       res.Id,
				Email:    res.Email,
				FullName: res.FullName,
				Access:   res.Access,
			},
		},
	})
}

func (r *AuthRouter) RegisterUser(c echo.Context) error {
	u := new(RegisterUserDTO)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(u); err != nil {
		msg := "Missing Field(s)"
		msgErr := err.Error()
		if strings.Contains("Email", msgErr) {
			msg = "Email not recognized"
		}
		return c.String(http.StatusBadRequest, msg)
	}
	res, err := r.controller.RegisterUser(&entity.User{
		Email:    u.Email,
		Password: u.Password,
		FullName: u.FullName,
	})

	if err != nil {
		if strings.ContainsAny(err.Error(), "duplicate") {
			return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
				Status: http.StatusBadRequest,
				Msg:    "Bad Request, Email not unique",
				Data:   nil,
			})
		}
		return c.String(http.StatusInternalServerError, "Bad request")
	}
	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "OK",
		Data: dto.UserInformationDTO{
			Id:       res.Id,
			Email:    res.Email,
			FullName: res.FullName,
			Access:   res.Access,
		},
	})
}
