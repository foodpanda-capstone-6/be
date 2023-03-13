package cart

import (
	"log"
	"vms-be/entities"
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
	log.Printf("[CartUseCase::AddToCart] un %s  qty %d mvId%d\n", username, qty, marketVoucherId)
	_ = uc.repos.Cart.Upsert(username, qty, marketVoucherId)
	return nil
}

func (uc *UseCase) GetCart(username string) []entities.VoucherInCart {
	log.Printf("[CartUseCase::GetCart] un %s  ", username)

	vcs, _ := uc.repos.Cart.GetByUsername(username)
	return vcs
}
