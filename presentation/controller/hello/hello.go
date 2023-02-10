package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type UseCaseHello interface {
	getHelloString() string
}

type Controller struct {
	usecase UseCaseHello
}

func (c *Controller) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/hello", c.hello)

	return r
}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	responseString := c.usecase.getHelloString()
	w.Write([]byte(responseString))
	return
}

func NewController() *Controller {
	return &Controller{}
}
