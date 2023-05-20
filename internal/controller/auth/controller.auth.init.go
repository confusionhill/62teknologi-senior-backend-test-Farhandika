package auth

import ar "project-template/internal/repository/auth"

type AuthController struct {
	repo *ar.AuthRepository
}

func Init(r *ar.AuthRepository) *AuthController {
	c := &AuthController{
		repo: r,
	}
	return c
}
