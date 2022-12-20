package engine

import "log"

type EngineOpts struct {
	LogPath string
	*DatabaseOpts
}
type Engine struct {
	DbC *DatabaseConnection
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
