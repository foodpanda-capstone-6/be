package incentive

import "vms-be/entities"

type UseCaseService interface {
	Commission(vcs []entities.VoucherInCart) error
	GetIncentivesOfUser(username string) []entities.Incentive
	Transfer(username, code string) error
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
