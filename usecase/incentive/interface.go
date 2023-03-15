package incentive

import "vms-be/entities"

type UseCaseService interface {
	Commission(vcs []entities.VoucherInCart) error
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
