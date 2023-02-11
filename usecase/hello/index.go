package hello

const HELLO_STRING = "hi from Voucher Management System."

type UseCase struct {
}

func (uc *UseCase) GetHelloString() string {
	return HELLO_STRING
}

func New() *UseCase {
	return &UseCase{}
}
