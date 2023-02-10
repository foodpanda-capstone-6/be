package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"vms-be/utils"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseOpts_SQL struct {
	Path string
}
type DatabaseOpts struct {
	DriverName string
	DatabaseOpts_SQL
}

type Database struct {
	*sql.DB
}

func ConnectSqlite3(path string) *Database {
	log.Println("[engine::database::connectSqlite3] connectSqlite3")

	fullpath, err := utils.GetFullPathOfPath(path)
	err = os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[engine::database::connectSqlite3] Creating Database at %s", fullpath)
	_, err = os.Create(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", fullpath)
	if err != nil {
		log.Fatalln("[Engine::ConnectDatabase::sqlite3] file cannot be opened.", err)
	} else {
		log.Printf("[Engine::ConnectDatabase::sqlite3] database at: %s", fullpath)
	}

	return &Database{DB: db}
}
