package main

import (
	"log"
	"os"
	"vms-be/core/entities"
	database2 "vms-be/core/infra/database"
	presentation "vms-be/core/presentation"
	authUC "vms-be/core/usecase/auth"
	helloUC "vms-be/core/usecase/hello"

	globalLog "vms-be/globallog"
)

var logOpts = &globalLog.EngineOpts{
	LogPath: "logs/engine.txt",
}

var DatabaseOpts = &database2.DatabaseOpts{DriverName: "sqlite3", DatabaseOpts_SQL: database2.DatabaseOpts_SQL{
	Path: "storage/main.db",
}}

var ServerConfig = &presentation.Opts{}

var GlobalLog *globalLog.GlobalLog

var MainCommand = struct {
	RunServer string
	SmokeTest string
}{
	RunServer: "run-server",
	SmokeTest: "smoke-test",
}

func DevUsers() []entities.LoginFields {
	return []entities.LoginFields{{Username: "kai", Password: "ilovepanda"}, {Username: "noel", Password: "pandaforlife"}, {Username: "naz", Password: "panda4ever"}}
}
func init() {

	args := os.Args
	var mainCommand string
	argsLength := len(args)
	if argsLength == 1 {
		mainCommand = MainCommand.RunServer
		log.Printf("[package::init] command defaulted to: %s", mainCommand)

	} else if argsLength > 1 {
		mainCommand = args[1]
		log.Printf("[package::init] command: %s", mainCommand)

	}
	switch mainCommand {
	case MainCommand.SmokeTest:
		GlobalLog = globalLog.InitGlobalLog(logOpts)
		os.Exit(0)
	case MainCommand.RunServer:
		GlobalLog = globalLog.InitGlobalLog(logOpts)

		ServerConfig.Addr = GetServerIngressPort()
		ServerConfig.LogPath = "./logs/log-server.txt"
		log.Printf("[package::init] Server Address: %s", ServerConfig.Addr)

		ServerConfig.ControllerArgs.Hello.UseCase = helloUC.New()

		db, err := database2.GetRepo(*DatabaseOpts)
		uc_auth := authUC.New(authUC.Args{Repos: authUC.Repos{Auth: db}})

		for _, loginFields := range DevUsers() {
			uc_auth.Register(loginFields)
		}

		ServerConfig.ControllerArgs.Auth.UseCase = uc_auth
		if err != nil {
			log.Fatalf("[InitEngine] db not initialize\n")
		}

		presentation.InitAndRunServer(ServerConfig)
	default:
		log.Printf("FAIL::[package::init] unknown command: %s", mainCommand)
		os.Exit(0)
	}

}
func main() {

}
