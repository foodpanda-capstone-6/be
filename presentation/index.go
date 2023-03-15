package server

import (
	"log"
	"net/http"
	controllerAuth "vms-be/presentation/controller/auth"
	controllerCart "vms-be/presentation/controller/cart"
	controllerHello "vms-be/presentation/controller/hello"
	controllerIncentive "vms-be/presentation/controller/incentives"
	controllerMarket "vms-be/presentation/controller/market"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"vms-be/presentation/middlewares"
)

type Opts struct {
	Addr           string
	LogPath        string
	ControllerArgs struct {
		Hello     controllerHello.Args
		Auth      controllerAuth.Args
		Market    controllerMarket.Args
		Cart      controllerCart.Args
		Incentive controllerIncentive.Args
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
	r.Route("/market", controllerMarket.New(opts.ControllerArgs.Market).Routes)
	r.Route("/cart", controllerCart.New(opts.ControllerArgs.Cart).Routes)
	r.Route("/incentives", controllerIncentive.New(opts.ControllerArgs.Incentive).Routes)
	// staticRouter := controllerStatic.New(STATIC_FOLDER)
	// r.Mount("/", staticRouter)

	http.ListenAndServe(opts.Addr, r)
}
