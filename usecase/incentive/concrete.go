package incentive

import (
	"fmt"
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
	code := fmt.Sprintf("PANDA%5d", globalIncentiveCodeSuffix)
	globalIncentiveCodeSuffix++

	return code
}

var globalTransferCodeSuffix int = 0

func generateTransferCode() string {
	code := fmt.Sprintf("TX%5d", globalTransferCodeSuffix)
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
	// to do

	return nil
}
