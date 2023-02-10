package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vms-be/utils"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseOpts_SQL struct {
	Path string
}

type DatabaseSQLite3 struct {
	*sql.DB
}

func (db *DatabaseSQLite3) Login(email, passwordHashed string) (bool, error) {
	err := db.QueryRow("SELECT `email` from `user` where email=? and hashed_password=?",
		email, passwordHashed).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("engine.login: credentials mismatched for email %s", email)
		}
		return false, fmt.Errorf("engine.login general database error")
	}

	return true, nil
}

func UseSqlite3(path string) *DatabaseSQLite3 {
	log.Println("[INFRA::connectSqlite3] connectSqlite3")

	fullpath, err := utils.GetFullPathOfPath(path)
	err = os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[INFRA::connectSqlite3] Creating Database at %s", fullpath)
	_, err = os.Create(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", fullpath)
	if err != nil {
		log.Fatalln("[INFRA::sqlite3] file cannot be opened.", err)
	} else {
		log.Printf("[INFRA::sqlite3] database at: %s", fullpath)
	}

	return &DatabaseSQLite3{DB: db}
}
