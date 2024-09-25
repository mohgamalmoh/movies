package auth

import (
	"movies/definitions/users"
)

type AuthService interface {
	GetAuthUser() users.User
}
