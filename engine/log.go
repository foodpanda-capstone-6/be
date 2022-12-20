package engine

import (
	"log"
	"os"
	"path/filepath"
	"time"
	"vms-be/utils"
)

func (m *Engine) InitLog(path string) {

	fullpath, err := utils.GetFullPathOfPath(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Setting log file path to: %s", fullpath)

	err = os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Printf("[Main:::InitLog] log file set to %s", fullpath)
	log.Printf("[Main:::InitLog] Time: %s", time.Now().String())

}
