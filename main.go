package main

import (
	"log"
	"os"

	"vms-be/engine"
)

var engineOpt = &engine.EngineOpts{
	LogPath: "logs/log.txt",
	DatabaseOpts: &engine.DatabaseOpts{DriverName: "sqlite3", DatabaseOpts_SQL: engine.DatabaseOpts_SQL{
		Path: "storage/main.db",
	}},
}

var globalEngine *engine.Engine

var MAIN_COMMAND = struct {
	RUN_SERVER string
	SMOKE_TEST string
}{
	RUN_SERVER: "run-server",
	SMOKE_TEST: "smoke-test",
}

func main() {

	args := os.Args
	var main_command string
	args_length := len(args)
	if args_length == 1 {
		main_command = MAIN_COMMAND.RUN_SERVER
		log.Printf("[main] command defaulted to: %s", main_command)

	} else if args_length > 1 {
		main_command = args[1]
		log.Printf("[main] command: %s", main_command)

	}
	switch main_command {
	case MAIN_COMMAND.SMOKE_TEST:
		globalEngine = engine.InitEngine(engineOpt)
		os.Exit(0)
	case MAIN_COMMAND.RUN_SERVER:
		globalEngine = engine.InitEngine(engineOpt)
	default:
		log.Printf("FAIL::[main] unknown command: %s", main_command)
		os.Exit(0)
	}

	if globalEngine != nil {
		defer globalEngine.TearDown()
	}

}
