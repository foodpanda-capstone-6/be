package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	controllerAuth "vms-be/core/presentation/controller/auth"
	controllerHello "vms-be/core/presentation/controller/hello"
	"vms-be/core/presentation/middlewares"
)

type Opts struct {
	Addr           string
	LogPath        string
	ControllerArgs struct {
		Hello controllerHello.Args
		Auth  controllerAuth.Args
	}
}

type Presentation struct {
	Opts
}

const STATIC_FOLDER = "./dist"

func InitAndRunServer(opts *Opts) {
	log.Println("[RunServer]")

	r := chi.NewRouter()

	r.Use(middlewares.NewLogger(opts.LogPath))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/auth", controllerAuth.New(opts.ControllerArgs.Auth).Routes)
	r.Route("/hello", controllerHello.New(opts.ControllerArgs.Hello).Routes)

	// staticRouter := controllerStatic.New(STATIC_FOLDER)
	// r.Mount("/", staticRouter)

	http.ListenAndServe(opts.Addr, r)
}
