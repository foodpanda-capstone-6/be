package main

import (
	"log"
	"os"
	"vms-be/entities"
	"vms-be/infra/database"
	inAuth "vms-be/infra/database/auth"
	inCart "vms-be/infra/database/cart"
	inIncentives "vms-be/infra/database/incentives"
	inMarket "vms-be/infra/database/market"
	presentation "vms-be/presentation"
	ucAuth "vms-be/usecase/auth"
	ucCart "vms-be/usecase/cart"
	ucHello "vms-be/usecase/hello"
	ucIncentive "vms-be/usecase/incentive"
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

		inAuth, err := inAuth.GetRepo(*DatabaseOpts)
		ucAuth := ucAuth.New(ucAuth.Args{Repos: ucAuth.Repos{Auth: inAuth}})

		inMarket, err := inMarket.GetRepo(*DatabaseOpts)
		ucMarket := ucMarket.New(ucMarket.Args{Repos: ucMarket.Repos{Market: inMarket}})

		inIncentives, err := inIncentives.GetRepo(*DatabaseOpts)
		ucIncentive := ucIncentive.New(ucIncentive.Args{Repos: ucIncentive.Repos{Incentives: inIncentives}})

		inCart, err := inCart.GetRepo(*DatabaseOpts)
		ucCart := ucCart.New(ucCart.Args{Repos: ucCart.Repos{Cart: inCart}, Services: ucCart.Services{UcIncentive: ucIncentive}})

		inAuth.Seed("schemas/users.sql")
		inMarket.Seed("schemas/market.sql")
		inIncentives.Seed("schemas/incentives.sql")
		inCart.Seed("schemas/cart.sql")

		for _, loginFields := range DevUsers() {
			ucAuth.Register(loginFields)
		}

		ServerConfig.ControllerArgs.Auth.UseCase = ucAuth
		ServerConfig.ControllerArgs.Market.UseCase = ucMarket
		ServerConfig.ControllerArgs.Cart.UseCase = ucCart
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
