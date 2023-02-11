package auth

import (
	"errors"
	"vms-be/entities"
	"vms-be/infra/database"
	auth "vms-be/usecase/auth/jwt"
)

type UseCase struct {
	repos    Repos
	services Services
}

type UseCaseInterface interface {
	Login(fields entities.LoginFields) (auth.JwtString, error)
}

var ErrLoginFailed error = errors.New("login failed")

func (uc *UseCase) Login(fields entities.LoginFields) (auth.JwtString, error) {
	username := fields.Username
	password := fields.Password
	ok, err := uc.repos.Auth.Login(username, password)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrLoginFailed
	}
	return uc.services.GenerateJWTString(username)
}

type Repos struct {
	Auth AuthRepo
}

type AuthRepo interface {
	database.InfraLoginService
}

type Services struct {
	auth.JWT
}
type Args struct {
	Repos
	Services
}

func New(args Args) *UseCase {
	return &UseCase{repos: args.Repos}
}
