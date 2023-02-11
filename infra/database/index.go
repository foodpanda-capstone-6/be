package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseOpts struct {
	DriverName string
	DatabaseOpts_SQL
}
type InfraService interface {
	InfraLoginService
}

type InfraLoginService interface {
	Login(username, password string) (bool, error)
}

func GetRepo(opts DatabaseOpts) (InfraService, error) {
	if opts.DriverName == "sqlite3" {
		return UseSqlite3(opts.Path), nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}