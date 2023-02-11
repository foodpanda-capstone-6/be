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
	Login(fields entities.LoginFields) (auth.String, error)
}

var ErrLoginFailed error = errors.New("login failed")

func hashPassword(password string) string {
	// TODO
	return password + "123"
}
func (uc *UseCase) Login(fields entities.LoginFields) (auth.String, error) {
	username := fields.Username
	hashedPassword := hashPassword(fields.Password)

	ok, err := uc.repos.Auth.Login(username, hashedPassword)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrLoginFailed
	}
	return uc.services.GenerateJWTString(username)
}

func (uc *UseCase) Register(fields entities.LoginFields) (auth.String, error) {
	username := fields.Username
	hashedPassword := hashPassword(fields.Password)

	ok, err := uc.repos.Auth.Register(username, hashedPassword)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrLoginFailed
	}
	return uc.services.GenerateJWTString(username)
}

type Repos struct {
	Auth Repo
}

type Repo interface {
	database.InfraAuthService
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
