package main

import (
	"log"
	"os"
	"vms-be/entities"
	"vms-be/infra/database"
	inAuth "vms-be/infra/database/auth"
	inMarket "vms-be/infra/database/market"
	presentation "vms-be/presentation"
	ucAuth "vms-be/usecase/auth"
	ucHello "vms-be/usecase/hello"
	ucMarket "vms-be/usecase/market"

	globalLog "vms-be/globallog"
)

var logOpts = &globalLog.EngineOpts{
	LogPath: "logs/engine.txt",
}

var DatabaseOpts = &database.Opts{DriverName: "sqlite3", OptsSQLite3: database.OptsSQLite3{
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

		ServerConfig.ControllerArgs.Hello.UseCase = ucHello.New()

		authInfra, err := inAuth.GetRepo(*DatabaseOpts)
		ucAuth := ucAuth.New(ucAuth.Args{Repos: ucAuth.Repos{Auth: authInfra}})

		inMarket, err := inMarket.GetRepo(*DatabaseOpts)
		ucMarket.New(ucMarket.Args{Repos: ucMarket.Repos{Market: inMarket}})
		authInfra.Seed("schemas/users.sql")
		inMarket.Seed("schemas/market.sql")
		for _, loginFields := range DevUsers() {
			ucAuth.Register(loginFields)
		}

		ServerConfig.ControllerArgs.Auth.UseCase = ucAuth
		if err != nil {
			log.Fatalf("[InitEngine] authInfra not initialize\n")
		}

		presentation.InitAndRunServer(ServerConfig)
	default:
		log.Printf("FAIL::[package::init] unknown command: %s", mainCommand)
		os.Exit(0)
	}

}
func main() {

}
