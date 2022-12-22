package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type ControllerHello struct {
}

func (c *ControllerHello) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/hello", c.hello)

	return r
}

var HELLO_STRING = "hi from Voucher Management System."

func (c *ControllerHello) hello(w http.ResponseWriter, r *http.Request) {

	log.Println("[ControllerHello.hello]")
	var HelloResponseString = HELLO_STRING
	w.Write([]byte(HelloResponseString))
	return

}
