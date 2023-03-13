package market

import (
	"log"
	"net/http"
	ucmarket "vms-be/usecase/market"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase ucmarket.UseCaseService
}

type Args struct {
	UseCase ucmarket.UseCaseService
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/check", c.hello)
	r.Post("/", c.login)

}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from auth"))
	return
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerAuth.hello]")

	return
}

func New(args Args) *Controller {
	return &Controller{useCase: args.UseCase}
}
