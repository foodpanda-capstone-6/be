package market

import (
	"vms-be/entities"
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

func (uc *UseCase) GetMarketVouchers() []entities.MarketVoucher {
	mvs, _ := uc.repos.Market.GetMarketVouchers()

	return mvs
}

type UseCaseService interface {
	GetMarketVouchers() []entities.MarketVoucher
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
