package cart

import "vms-be/entities"

type UseCaseService interface {
	AddToCart(username string, qty int, marketVoucherId int) error
	GetCart(username string) []entities.VoucherInCart
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
