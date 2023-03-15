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

	IncentiveGetter interface {
		GetIncentivesOfUser(username string) ([]entities.Incentive, error)
	}
)

type InfraService interface {
	InfraSeeder
	IncentiveCommissioner
	IncentiveGetter
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

func (db *RepoSQLite3) GetIncentivesOfUser(username string) ([]entities.Incentive, error) {

	log.Printf("[GetIncentivesOfUser] incentive: un %s \n", username)

	rows, err := db.Query("SELECT id, username, incentive_code, transfer_code, value from `incentives` where username=?", username)
	if err != nil {
		log.Printf("[GetIncentivesOfUser] get incentives error is after attempting to get :%s \n", err)
	}
	defer rows.Close()

	var Incentives = make([]entities.Incentive, 0)
	for rows.Next() {
		var Id int
		var Username string
		var IncentiveCode string
		var TransferCode string
		var Value int

		err = rows.Scan(&Id, &Username, &IncentiveCode, &TransferCode, &Value)

		if err != nil {
			log.Printf("[GetByUsername] scan cart entry error is after attempting to get %v \n", err)
		}
		Incentives = append(Incentives, entities.Incentive{Id: Id, Username: Username, IncentiveCode: IncentiveCode, TransferCode: TransferCode, Value: Value})
	}

	return Incentives, nil
}

func GetRepo(opts Opts) (InfraService, error) {

	if opts.DriverName == "sqlite3" {
		db := database.UseSqlite3(opts.Path)
		return &RepoSQLite3{DB: db}, nil
	}

	return nil, fmt.Errorf("[GetRepo] Invalid driver name %s \n", opts.DriverName)
}
