package incentives

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vms-be/entities"
	"vms-be/infra/database"
	"vms-be/infra/database/entries"
)

type Opts = database.Opts

type (
	InfraSeeder interface {
		Seed(schemaPath string)
	}

	IncentiveCommissioner interface {
		Commission(_ []entities.Incentive) error
	}
)

type InfraService interface {
	InfraSeeder
	IncentiveCommissioner
}

type RepoSQLite3 struct {
	*sql.DB
	InfraService
}

func (db *RepoSQLite3) Seed(schemaPath string) {
	log.Println("[Seeding incentives]")
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

func EntityToEntryIncentive(entities []entities.Incentive) []entries.Incentive {

	IncentiveEntries := make([]entries.Incentive, 0)
	for _, entity := range entities {
		IncentiveEntries = append(IncentiveEntries, entries.Incentive{
			IncentiveCode: entity.IncentiveCode,
			Value:         entity.Value,
			TransferCode:  entity.TransferCode,
			Username:      entity.Username,
		})
	}

	return IncentiveEntries
}

func (db *RepoSQLite3) Commission(ins []entities.Incentive) error {

	log.Printf("[Commission] incentives: un %v \n", ins)

	incentiveEntries := EntityToEntryIncentive(ins)
	for _, in := range incentiveEntries {
		_ = db.QueryRow("INSERT INTO incentives (username, incentive_code, transfer_code, value) VALUES (?,?,?,?)", in.Username, in.IncentiveCode, in.TransferCode, in.Value).Scan()
	}

	return nil
}

func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
