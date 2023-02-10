package engine

import (
	"log"

	"vms-be/infra/database"
)

type EngineOpts struct {
	LogPath string
	*database.DatabaseOpts
}

type DatabaseService interface {
	closeDB()
}
type Engine struct {
	services struct {
		dbService DatabaseService
	}
}

func InitEngine(opts *EngineOpts) *Engine {

	log.Println("[InitEngine]")
	engine := &Engine{}

	engine.InitLog(opts.LogPath)

	return engine
}
