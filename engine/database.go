package engine

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

type DatabaseConnection struct {
	DB *sql.DB
}

func (dbC *DatabaseConnection) Close() {
	dbC.DB.Close()
}

func connectSqlite3(path string) *sql.DB {
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

	return db
}

func (e *Engine) ConnectDatabase(opts *DatabaseOpts) {
	e.DbC = &DatabaseConnection{}
	log.Println("[ConnectDatabase]")
	if opts.DriverName == "sqlite3" {
		e.DbC.DB = connectSqlite3(opts.Path)
	} else {
		log.Fatalln("[Engine::ConnectDatabase] mode not specified")
	}
}
