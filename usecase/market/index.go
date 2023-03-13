package market

import (
	marketAuth "vms-be/infra/database/market"
	auth "vms-be/usecase/auth/jwt"
)

type Repos struct {
	Market marketAuth.InfraService
}

type Services struct {
	auth.JWT
}

type UseCase struct {
	repos    Repos
	services Services
}

type UseCaseService interface {
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
