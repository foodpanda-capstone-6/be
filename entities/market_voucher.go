package entities

type MarketVoucher struct {
	Id          int
	Description string
	Amount      int
}

type VoucherInCart struct {
	Username string
	Id       int
	Qty      int
	Amount   int
}
