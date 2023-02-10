package server

import (
	"log"
	"net/http"
	controllerHello "vms-be/presentation/controller/hello"
	controllerStatic "vms-be/presentation/controller/static"
	middlewares "vms-be/presentation/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Opts struct {
	Addr    string
	LogPath string
}

type Presentation struct {
	Opts
}

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

	staticRouter := controllerStatic.NewController("./dist")
	r.Mount("/", staticRouter)

	r.Mount("/hello", controllerHello.NewController().Routes())
	http.ListenAndServe(opts.Addr, r)
}
