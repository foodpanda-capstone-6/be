package engine

import (
	"database/sql"
	"log"

	"vms-be/engine/database"
)

type EngineOpts struct {
	LogPath string
	*database.DatabaseOpts
}
type Engine struct {
	DbC *sql.DB
}

func InitEngine(opts *EngineOpts) *Engine {

	log.Println("[InitEngine]")
	engine := &Engine{}

	engine.InitLog(opts.LogPath)

	engine.ConnectDatabase(opts.DatabaseOpts)
	return engine
}

func (engine *Engine) closeDB() {
	log.Println("[engine::closeDB]")
	engine.DbC.Close()
}

func (engine *Engine) TearDown() {
	log.Println("[engine::TearDown]")
	go engine.closeDB()
}

func (e *Engine) ConnectDatabase(opts *database.DatabaseOpts) {
	log.Println("[ConnectDatabase]")
	if opts.DriverName == "sqlite3" {
		e.DbC = database.ConnectSqlite3(opts.Path)
	} else {
		log.Fatalln("[Engine::ConnectDatabase] mode not specified")
	}
}
