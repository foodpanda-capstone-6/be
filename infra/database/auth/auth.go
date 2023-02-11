package auth

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
	InfraLoginAuthenticator interface {
		Login(username, hashedPassword string) (bool, error)
	}
	InfraUserRegister interface {
		Register(username, passwordHashed string) (bool, error)
	}
	InfraSeeder interface {
		Seed(schemaPath string)
	}
)

type InfraService interface {
	InfraLoginAuthenticator
	InfraUserRegister
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
func (db *RepoSQLite3) Login(username, passwordHashed string) (bool, error) {

	log.Printf("[Login] username: %s \n", username)
	var ResultUsername string
	err := db.QueryRow("SELECT `username` from `users` where username=? and hashed_password=?", username, passwordHashed).Scan(&ResultUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[Login] FAIL  %v \n", err)
			return false, fmt.Errorf("[Login]: credentials mismatched for username %s", username)
		}
		log.Printf("[Login] FAIL  %v \n", err)
		return false, fmt.Errorf("[Login] general database error")
	}

	return true, nil
}

func (db *RepoSQLite3) Register(username, passwordHashed string) (bool, error) {

	log.Printf("[Infra::Register] username: %s \n", username)
	err := db.QueryRow("INSERT INTO `users` (username,hashed_password) VALUES (?,?)", username, passwordHashed).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[Infra::Register] FAIL  %v \n", err)
			return false, fmt.Errorf("[Infra::Register]: credentials mismatched for email %s", username)
		}
		log.Printf("[Infra::Register] FAIL  %v \n", err)
		return false, fmt.Errorf("[Infra::Register] general database error")
	}

	return true, nil
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
