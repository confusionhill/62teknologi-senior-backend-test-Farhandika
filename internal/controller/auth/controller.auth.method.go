package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"project-template/model/entity"
)

func (c *AuthController) RegisterUser(newUser *entity.User) (*entity.User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return nil, err
	}
	newUser.Password = string(bytes)
	return c.repo.InsertNewUser(newUser)
}

func (c *AuthController) LoginUser(user *entity.User) (*entity.User, error) {
	result, err := c.repo.GetUserInfoByEmail(user.Email)
	if err != nil {
		return nil, errors.New("Cannot find user")
	}
	isPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if isPass != nil {
		return nil, errors.New("Password does not match")
	}
	return result, err
}
