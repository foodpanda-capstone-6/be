package incentive

import (
	"fmt"
	"log"
	"math/rand"
	"vms-be/entities"
	incentives "vms-be/infra/database/incentives"
)

type Repos struct {
	Incentives incentives.InfraService
}

type Services struct {
}

type UseCase struct {
	repos    Repos
	services Services
}

var globalIncentiveCodeSuffix int = 0

func generateIncentiveCode() string {
	code := fmt.Sprintf("PANDA%05d", globalIncentiveCodeSuffix)
	globalIncentiveCodeSuffix++

	return code
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(randSeq(10))
}

var globalTransferCodeSuffix int = 0

func generateTransferCode() string {
	code := fmt.Sprintf("%s%05d", randSeq(5), globalTransferCodeSuffix)
	globalTransferCodeSuffix++
	return code
}

func VoucherInCartToIncentive(incentive entities.VoucherInCart) entities.Incentive {
	return entities.Incentive{
		IncentiveCode: generateIncentiveCode(),
		TransferCode:  generateTransferCode(),
		Username:      incentive.Username,
		Value:         incentive.Amount,
	}
}

func VouchersInCartToIncentives(vcs []entities.VoucherInCart) []entities.Incentive {
	ins := make([]entities.Incentive, 0)

	for _, vc := range vcs {
		ins = append(ins, VoucherInCartToIncentive(vc))
	}

	return ins
}

func (uc *UseCase) Commission(vcs []entities.VoucherInCart) error {
	uc.repos.Incentives.Commission(VouchersInCartToIncentives(vcs))
	return nil
}

func (uc *UseCase) GetIncentivesOfUser(username string) []entities.Incentive {

	ins, err := uc.repos.Incentives.GetIncentivesOfUser(username)

	if err != nil {
		log.Printf("IncentiveUseCase: Error getting Incentives %s \n", err.Error())
	} else {
		log.Printf("IncentiveUseCase: ok getting Incentives %v \n", ins)
	}
	return ins
}
