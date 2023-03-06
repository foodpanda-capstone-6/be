package hello

const HELLO_STRING = "hi from Voucher Management System."

type UseCase struct {
}

type Service interface {
	GetHelloString() string
}

func (uc *UseCase) GetHelloString() string {
	return HELLO_STRING
}

func New() Service {
	return &UseCase{}
}
