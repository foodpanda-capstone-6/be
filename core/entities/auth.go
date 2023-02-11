package entities

type LoginFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLoginFields(username, password string) LoginFields {
	return LoginFields{Username: username, Password: password}
}
