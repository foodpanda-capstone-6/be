package main

import (
	"log"
	"os"

	"vms-be/engine"
	"vms-be/infra/database"
	presentation "vms-be/presentation"
)

var engineOpt = &engine.EngineOpts{
	LogPath: "logs/engine.txt",
	DatabaseOpts: &database.DatabaseOpts{DriverName: "sqlites3", DatabaseOpts_SQL: database.DatabaseOpts_SQL{
		Path: "storage/main.db",
	}},
}

var ServerConfig = &presentation.Opts{}

var globalEngine *engine.Engine

var MAIN_COMMAND = struct {
	RUN_SERVER string
	SMOKE_TEST string
}{
	RUN_SERVER: "run-server",
	SMOKE_TEST: "smoke-test",
}

func init() {

	args := os.Args
	var main_command string
	args_length := len(args)
	if args_length == 1 {
		main_command = MAIN_COMMAND.RUN_SERVER
		log.Printf("[package::init] command defaulted to: %s", main_command)

	} else if args_length > 1 {
		main_command = args[1]
		log.Printf("[package::init] command: %s", main_command)

	}
	switch main_command {
	case MAIN_COMMAND.SMOKE_TEST:
		globalEngine = engine.InitEngine(engineOpt)
		os.Exit(0)
	case MAIN_COMMAND.RUN_SERVER:
		globalEngine = engine.InitEngine(engineOpt)
		ServerConfig.Addr = GetServerIngressPort()
		ServerConfig.LogPath = "./logs/log-server.txt"
		log.Printf("[package::init] Server Address: %s", ServerConfig.Addr)
		presentation.InitAndRunServer(ServerConfig)
	default:
		log.Printf("FAIL::[package::init] unknown command: %s", main_command)
		os.Exit(0)
	}

}
func main() {

	if globalEngine != nil {
		defer globalEngine.TearDown()
	}

}
