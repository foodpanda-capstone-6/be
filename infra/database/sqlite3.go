package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"vms-be/utils"

	_ "github.com/mattn/go-sqlite3"
)

type OptsSQLite3 struct {
	Path string
}

type SQLite3 struct {
	*sql.DB
	InfraAuthService
}

func (db *SQLite3) Login(username, passwordHashed string) (bool, error) {

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

func (db *SQLite3) Register(username, passwordHashed string) (bool, error) {

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

func UseSqlite3(path string) *SQLite3 {
	log.Println("[INFRA::connectSqlite3] connectSqlite3")
	// OPEN
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

	// CREATE TABLES

	pathSchema := filepath.Join("schema.sql")

	c, err := ioutil.ReadFile(pathSchema)
	if err != nil {
		log.Fatal(err)
	}
	sql := string(c)
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	return &SQLite3{DB: db}
}
