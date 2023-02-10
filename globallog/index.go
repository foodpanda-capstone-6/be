package engine

import (
	"log"

	"vms-be/infra/database"
)

type EngineOpts struct {
	LogPath string
	*database.DatabaseOpts
}

type GlobalLog struct {
}

func InitGlobalLog(opts *EngineOpts) *GlobalLog {

	log.Println("[InitEngine]")
	GlobalLog := &GlobalLog{}

	GlobalLog.InitLog(opts.LogPath)

	return GlobalLog
}
