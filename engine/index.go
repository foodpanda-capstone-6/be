package engine

import (
	"log"

	"vms-be/infra/database"
)

type EngineOpts struct {
	LogPath string
	*database.DatabaseOpts
}

type Engine struct {
	Services struct {
		dbService database.InfraService
	}
}

func (e *Engine) TearDown() {

}

func InitEngine(opts *EngineOpts) *Engine {

	log.Println("[InitEngine]")
	engine := &Engine{}

	engine.InitLog(opts.LogPath)
	engine.Services.dbService = database.UseSqlite3(opts.DatabaseOpts.Path)
	return engine
}
