package cart

import (
	"log"
	"vms-be/entities"
	cart "vms-be/infra/database/cart"
	auth "vms-be/usecase/auth/jwt"
	incentive "vms-be/usecase/incentive"
	ucMarket "vms-be/usecase/market"
)

type Repos struct {
	Cart cart.InfraService
}

type Services struct {
	auth.JWT
	ucMarket    ucMarket.UseCaseService
	UcIncentive incentive.UseCaseService
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

func (uc *UseCase) makePayment(vcs []entities.VoucherInCart) []entities.VoucherInCart {
	// to do

	return vcs
}

func (uc *UseCase) removeVouchers(vcs []entities.VoucherInCart) {
	// to do

}

func (uc *UseCase) Purchase(username string) {
	log.Printf("[CartUseCase::Purchase] un %s  ", username)

	vcs, _ := uc.repos.Cart.GetByUsername(username)

	uc.makePayment(vcs)
	uc.services.UcIncentive.Commission(vcs)
	uc.removeVouchers(vcs)

}
