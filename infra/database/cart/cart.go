package cart

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vms-be/infra/database"
)

type Opts = database.Opts

type (
	InfraSeeder interface {
		Seed(schemaPath string)
	}

	CartAdder interface {
		Upsert(username string, quantity int, marketVoucherId int) error
	}
)

type InfraService interface {
	InfraSeeder
	CartAdder
}

type RepoSQLite3 struct {
	*sql.DB
	InfraService
}

func (db *RepoSQLite3) Seed(schemaPath string) {
	log.Println("[Seeding cart]")
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

func (db *RepoSQLite3) Upsert(username string, quantity int, marketVoucherId int) error {
	return nil
}

func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
