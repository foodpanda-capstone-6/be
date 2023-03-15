package cart

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

	CartAdder interface {
		Upsert(username string, quantity int, marketVoucherId int) error
	}

	CartGetter interface {
		GetByUsername(username string) ([]entities.VoucherInCart, error)
	}
)

type InfraService interface {
	InfraSeeder
	CartAdder
	CartGetter
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
	log.Printf("[Upsert] cart: un %s \n", username)
	var ResultUsername string

	err := db.QueryRow("SELECT username from `cart` where username=? and market_voucher_id=?", username, marketVoucherId).Scan(&ResultUsername)

	if err == sql.ErrNoRows { // entry not exist
		log.Printf("[Upsert] cart entry empty. attempting to insert %v \n", err)
		_ = db.QueryRow("INSERT INTO cart (username, market_voucher_id, qty) VALUES (?,?,?)", username, marketVoucherId, quantity).Scan()
		return nil
	}

	if err != nil {

		log.Printf("[Upsert] FAIL  %v \n", err)
		return errors.New("")
	}
	err = db.QueryRow("UPDATE cart set qty=? where username=? and market_voucher_id=?", quantity, username, marketVoucherId).Scan()
	log.Printf("[Upsert] cart entry exist. after attempting to update %v \n", err)

	return nil
}
func (db *RepoSQLite3) GetByUsername(username string) ([]entities.VoucherInCart, error) {
	log.Printf("[GetByUsername] cart: un %s \n", username)

	rows, err := db.Query("SELECT username, market_voucher_id, qty, amount from `cart`,`market` where username=? and `cart`.`market_voucher_id` = `market`.`id`", username)
	defer rows.Close()

	if err != nil {
		log.Printf("[GetByUsername] cart error is after attempting to get %s \n", err)
	}

	var CartVouchers = make([]entities.VoucherInCart, 0)
	for rows.Next() {

		var Username string
		var MarketVoucherId int
		var Amount int
		var Qty int

		err = rows.Scan(&Username, &MarketVoucherId, &Qty, &Amount)

		if err != nil {
			log.Printf("[GetByUsername] scan cart entry error is after attempting to get %v \n", err)
		}
		CartVouchers = append(CartVouchers, entities.VoucherInCart{Username: Username, Id: MarketVoucherId, Qty: Qty, Amount: Amount})
	}

	return CartVouchers, nil
}

func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
