package hello

import (
	"log"
	"net/http"
	"vms-be/usecase/hello"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase hello.Service
}

type Args struct {
	UseCase hello.Service
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/", c.hello)
	r.Get("/hello", c.helloSub)
	r.Get("/hello/", c.helloSubSlash)

}
func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	responseString := c.useCase.GetHelloString()
	w.Write([]byte(responseString))
	return
}

func (c *Controller) helloSub(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	responseString := c.useCase.GetHelloString()
	responseString = "helloSub"
	w.Write([]byte(responseString))
	return
}

func (c *Controller) helloSubSlash(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerHello.hello]")
	_ = c.useCase.GetHelloString()
	w.Write([]byte("subsl"))
	return
}

func New(arg Args) *Controller {
	return &Controller{useCase: arg.UseCase}
}
