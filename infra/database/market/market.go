package market

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

import "vms-be/infra/database"

type Opts = database.Opts

type (
	InfraSeeder interface {
		Seed(schemaPath string)
	}
)

type InfraService interface {
	InfraSeeder
}

type RepoSQLite3 struct {
	*sql.DB
	InfraService
}

func (db *RepoSQLite3) Seed(schemaPath string) {

	pathSchema := filepath.Join(schemaPath)
	c, err := os.ReadFile(pathSchema)
	if err != nil {
		log.Fatal(err)
	}
	sql := string(c)
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRepo(opts Opts) (InfraService, error) {

	var infra InfraService
	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		Repo := &RepoSQLite3{DB: db}
		infra = Repo

		return infra, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
