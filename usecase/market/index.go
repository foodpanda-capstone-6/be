package market

type UseCase struct {
}

type Args struct{}

func New(args Args) *UseCase {
	return &UseCase{}
}
