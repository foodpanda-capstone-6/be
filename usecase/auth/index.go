package auth

import (
	"errors"
	"vms-be/entities"
	auth "vms-be/usecase/auth/jwt"
)

type UseCase struct {
}

type UseCaseInterface interface {
	Login(fields entities.LoginFields) (auth.JwtString, error)
}

func (uc *UseCase) Login(fields entities.LoginFields) (auth.JwtString, error) {
	return "", errors.New("login failed")
}

type Repos struct {
	Auth AuthRepo
}

type AuthRepo interface {
	Login(username, password string) (bool, error)
}

type Services struct {
	auth.JWT
}
type Args struct {
	Repos
	Services
}

func New(args Args) *UseCase {
	return &UseCase{}
}
