package middlewares

import (
	"log"
	"net/http"
	"vms-be/utils"

	"github.com/go-chi/chi/v5/middleware"
)

func NewLogger(path string) func(http.Handler) http.Handler {

	fullpath, logWriter, err := utils.OpenOrCreateFile(path)
	log.Printf("[Server.middlewares.NewLogger] Setting file path to: %s", fullpath)

	if err != nil {
		log.Fatalln("FAIL::[RunServer] Fail to create new logger.", err)
	} else {
		log.Printf("Logger created for server %s", path)
	}

	mw := &middleware.DefaultLogFormatter{Logger: log.New(logWriter, "", log.LstdFlags), NoColor: true}
	mw.Logger.Print("---------------------------------------------START---------------------------------------------------------------")
	return middleware.RequestLogger(mw)
}
