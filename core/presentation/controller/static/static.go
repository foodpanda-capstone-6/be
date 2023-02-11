package routes

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func New(path string) chi.Router {
	r := chi.NewRouter()

	staticPath, err := filepath.Abs(path)

	if err != nil {
		log.Fatalln("FAIL::[routeStatic]", staticPath)
	}

	log.Printf("[routeStatic] Static Directory: %s", staticPath)
	fs := http.FileServer(http.Dir(staticPath))
	r.Handle("/*", http.StripPrefix("/", fs))

	return r
}
