package engine

import (
	"log"
	"time"
	"vms-be/utils"
)

func (m *GlobalLog) InitLog(path string) {
	fullpath, file, err := utils.OpenOrCreateFile(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[GlobalLog::InitLog] NOTE: setting log file to %s. Redirecting logs... ", fullpath)

	log.SetOutput(file)

	log.Printf("[GlobalLog::InitLog] log file set to %s", fullpath)
	log.Println("---------------------------------------------START---------------------------------------------------------------")
	log.Printf("[GlobalLog::InitLog] Time: %s", time.Now().String())
}
