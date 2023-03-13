package cart

type UseCaseService interface {
	AddToCart(username string, qty int, marketVoucherId int) error
}

type Args struct {
	Repos
	Services
}

func New(args Args) UseCaseService {
	return &UseCase{repos: args.Repos, services: args.Services}
}
