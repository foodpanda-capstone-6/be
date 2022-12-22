package engine

import (
	"log"
	"time"
	"vms-be/utils"
)

func (m *Engine) InitLog(path string) {

	fullpath, file, err := utils.OpenOrCreateFile(path)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Printf("[Engine:::InitLog] log file set to %s", fullpath)
	log.Println("---------------------------------------------START---------------------------------------------------------------")
	log.Printf("[Engine:::InitLog] Time: %s", time.Now().String())
}
