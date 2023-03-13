package market

import (
	"database/sql"
	"errors"
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
	// IMPLEMENTME

	log.Printf("[MarketRepo] GetMarketVouchers \n")
	var MarketVouchers []entities.MarketVoucher
	// err := db.QueryRow("SELECT username, market_voucher_id, qty, amount from `cart`, `market` where `cart`.`username` = ? and `cart`.`market_voucher_id`=`market`.`id`", username).Scan(&MarketVouchers)

	rows, err := db.Query("SELECT id, description, amount from `market`")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[MarketRepo] FAIL no rows  %v \n", err)
			return []entities.MarketVoucher{}, errors.New("")
		}
		log.Printf("[MarketRepo] FAIL  %v \n", err)
		return []entities.MarketVoucher{}, fmt.Errorf("[MarketRepo] general database error")
	}

	defer rows.Close()

	for rows.Next() {
		var Id int
		var Description string
		var Amount int

		err = rows.Scan(&Id, &Description, &Amount)
		MarketVouchers = append(MarketVouchers, entities.MarketVoucher{Id: Id, Description: Description, Amount: Amount})
	}

	// return []entities.MarketVoucher{{Id: 999, Description: "hello", Amount: 100}}, nil
	return MarketVouchers, nil
}
func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
