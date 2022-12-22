package routes

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
)

func NewRouterStatic(path string) chi.Router {
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
