package server

import (
	"log"
	"net/http"
	"vms-be/server/middlewares"
	"vms-be/server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type ServerOpts struct {
	Addr string
}

type Endpoints struct {
}

func RunServer(opts *ServerOpts) {

	log.Println("[RunServer]")
	r := chi.NewRouter()

	r.Use(middlewares.NewLogger("./logs/log-server.txt"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	staticRouter := routes.NewRouterStatic("./dist")
	r.Mount("/", staticRouter)

	r.Mount("/hello", (&routes.ControllerHello{}).Routes())
	http.ListenAndServe(opts.Addr, r)
}
