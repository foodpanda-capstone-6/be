package hello

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UseCase interface {
	GetHelloString() string
}

type Controller struct {
	usecase UseCase
}

type Args struct {
	UseCase
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/", c.hello)
	r.Get("/hello", c.helloSub)
	r.Get("/hello/", c.helloSubSlash)

}
func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	responseString := c.usecase.GetHelloString()
	w.Write([]byte(responseString))
	return
}

func (c *Controller) helloSub(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	responseString := c.usecase.GetHelloString()
	responseString = "helloSub"
	w.Write([]byte(responseString))
	return
}

func (c *Controller) helloSubSlash(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	_ = c.usecase.GetHelloString()
	w.Write([]byte("subsl"))
	return
}

func New(arg Args) *Controller {
	return &Controller{usecase: arg.UseCase}
}
