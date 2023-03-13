package cart

import (
	cart "vms-be/infra/database/cart"
	auth "vms-be/usecase/auth/jwt"
	ucMarket "vms-be/usecase/market"
)

type Repos struct {
	Cart cart.InfraService
}

type Services struct {
	auth.JWT
	ucMarket.UseCaseService
}

type UseCase struct {
	repos    Repos
	services Services
}

func (uc *UseCase) AddToCart(username string, qty int, marketVoucherId int) error {
	return nil
}
