package hello

const HELLO_STRING = "hi from Voucher Management System."

type UseCase struct {
}

func (uc *UseCase) getHelloString() string {
	return HELLO_STRING
}
