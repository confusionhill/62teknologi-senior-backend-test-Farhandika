package auth

import (
	"project-template/model/entity"
)

func (r *AuthRepository) GetUserInfoByEmail(email string) (*entity.User, error) {
	return nil, nil
}

func (r *AuthRepository) InsertNewUser(newUser *entity.User) (*entity.User, error) {
	return newUser, nil
}
