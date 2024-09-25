package auth

import (
	"movies/definitions/auth"
	"movies/definitions/users"
)

type AuthServiceImpl struct {
}

func NewAuthServiceImpl() auth.AuthService {
	return &AuthServiceImpl{}
}

func (t *AuthServiceImpl) GetAuthUser() users.User {
	return users.User{
		Id:    1,
		Name:  "jimmi",
		Email: "mohgamalmoh",
	}
}
