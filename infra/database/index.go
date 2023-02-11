package database

import (
	_ "github.com/mattn/go-sqlite3"
)

type Opts struct {
	DriverName string
	OptsSQLite3
}
