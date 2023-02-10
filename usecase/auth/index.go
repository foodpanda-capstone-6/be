package auth

type UseCase struct {
}

func (uc *UseCase) Login() {}

type Repos struct {
	Auth AuthRepo
}

type AuthRepo interface {
	Login(username, password string) (bool, error)
}

func New(r Repos) *UseCase {
	return nil
}
