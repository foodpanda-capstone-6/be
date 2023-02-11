package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Opts struct {
	DriverName string
	OptsSQLite3
}
type InfraAuthService interface {
	InfraLoginAuthenticator
	InfraUserRegister
}

type (
	InfraLoginAuthenticator interface {
		Login(username, hashedPassword string) (bool, error)
	}
	InfraUserRegister interface {
		Register(username, passwordHashed string) (bool, error)
	}
)

func GetRepo(opts Opts) (InfraAuthService, error) {
	if opts.DriverName == "sqlite3" {
		return UseSqlite3(opts.Path), nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
