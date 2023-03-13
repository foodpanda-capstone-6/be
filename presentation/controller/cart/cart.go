package market

import (
	"log"
	"net/http"
	uccart "vms-be/usecase/cart"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase uccart.UseCaseService
}

func (c *Controller) Routes(r chi.Router) {
	r.Post("/all", c.getAll)
	r.Post("/add", c.add)
}

func (c *Controller) getAll(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerCart.login]")
	return
}

func (c *Controller) add(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerCart.add]")
	w.WriteHeader(http.StatusOK)
	return
}

type Args struct {
	UseCase uccart.UseCaseService
}

func New(args Args) *Controller {
	return &Controller{useCase: args.UseCase}
}
