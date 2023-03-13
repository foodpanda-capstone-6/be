package market

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vms-be/entities"
	"vms-be/infra/database"
)

type Opts = database.Opts

type (
	InfraSeeder interface {
		Seed(schemaPath string)
	}

	MarketVoucherGetter interface {
		GetMarketVouchers() ([]entities.MarketVoucher, error)
	}
)

type InfraService interface {
	InfraSeeder
	MarketVoucherGetter
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
func (db *RepoSQLite3) GetMarketVouchers() ([]entities.MarketVoucher, error) {

	return []entities.MarketVoucher{{Id: 999, Description: "hello", Amount: 100}}, nil
}
func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
