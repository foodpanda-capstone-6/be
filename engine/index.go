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
	db, err := database.GetRepo(*opts.DatabaseOpts)
	if err != nil {
		log.Fatalf("[InitEngine] db not initialized\n")
	}
	engine.Services.dbService = db

	return engine
}
