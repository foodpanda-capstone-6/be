package main

import (
	"log"
	"os"
	"vms-be/core/entities"
	database2 "vms-be/core/infra/database"
	presentation "vms-be/core/presentation"
	uc_auth "vms-be/core/usecase/auth"
	uc_hello "vms-be/core/usecase/hello"

	globallog "vms-be/globallog"
)

var logOpts = &globallog.EngineOpts{
	LogPath: "logs/engine.txt",
}

var DatabaseOpts = &database2.DatabaseOpts{DriverName: "sqlite3", DatabaseOpts_SQL: database2.DatabaseOpts_SQL{
	Path: "storage/main.db",
}}

var ServerConfig = &presentation.Opts{}

var GlobalLog *globallog.GlobalLog

var MAIN_COMMAND = struct {
	RUN_SERVER string
	SMOKE_TEST string
}{
	RUN_SERVER: "run-server",
	SMOKE_TEST: "smoke-test",
}

func DevUsers() []entities.LoginFields {
	return []entities.LoginFields{{Username: "kai", Password: "ilovepanda"}, {Username: "noel", Password: "pandaforlife"}, {Username: "naz", Password: "panda4ever"}}
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
		GlobalLog = globallog.InitGlobalLog(logOpts)
		os.Exit(0)
	case MAIN_COMMAND.RUN_SERVER:
		GlobalLog = globallog.InitGlobalLog(logOpts)

		ServerConfig.Addr = GetServerIngressPort()
		ServerConfig.LogPath = "./logs/log-server.txt"
		log.Printf("[package::init] Server Address: %s", ServerConfig.Addr)

		ServerConfig.ControllerArgs.Hello.UseCase = uc_hello.New()

		db, err := database2.GetRepo(*DatabaseOpts)
		uc_auth := uc_auth.New(uc_auth.Args{Repos: uc_auth.Repos{Auth: db}})

		for _, loginFields := range DevUsers() {
			uc_auth.Register(loginFields)
		}

		ServerConfig.ControllerArgs.Auth.UseCase = uc_auth
		if err != nil {
			log.Fatalf("[InitEngine] db not initialize\n")
		}

		presentation.InitAndRunServer(ServerConfig)
	default:
		log.Printf("FAIL::[package::init] unknown command: %s", main_command)
		os.Exit(0)
	}

}
func main() {

}
