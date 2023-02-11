package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"vms-be/utils"

	_ "github.com/mattn/go-sqlite3"
)

type OptsSQLite3 struct {
	Path string
}

func UseSqlite3(path string) *sql.DB {
	log.Println("[INFRA::connectSqlite3] connectSqlite3")
	// OPEN
	fullPath, err := utils.GetFullPathOfPath(path)
	err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[INFRA::connectSqlite3] Creating Database at %s", fullPath)
	_, err = os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", fullPath)
	if err != nil {
		log.Fatalln("[INFRA::sqlite3] file cannot be opened.", err)
	} else {
		log.Printf("[INFRA::sqlite3] database at: %s", fullPath)
	}

	return db
}
