package engine

import (
	"log"
)

type EngineOpts struct {
	LogPath string
}

type GlobalLog struct {
}

func InitGlobalLog(opts *EngineOpts) *GlobalLog {

	log.Println("[InitEngine]")
	GlobalLog := &GlobalLog{}

	GlobalLog.InitLog(opts.LogPath)

	return GlobalLog
}
