package market

import (
	"encoding/json"
	"log"
	"net/http"
	ucincentives "vms-be/usecase/incentive"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase ucincentives.UseCaseService
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/all", c.getIncentivesOfUser)
	r.Post("/transfer", c.transfer)
}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from incentives"))
	return
}

func (c *Controller) transfer(w http.ResponseWriter, r *http.Request) {

	log.Println("[transfer]")
	code := r.URL.Query().Get("code")
	username := r.URL.Query().Get("username")
	err := c.useCase.Transfer(username, code)

	if err != nil {
		log.Printf("[transfer] error %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (c *Controller) getIncentivesOfUser(w http.ResponseWriter, r *http.Request) {

	log.Println("[getIncentivesOfUser]")
	username := r.URL.Query().Get("username")
	ins := c.useCase.GetIncentivesOfUser(username)
	responseString, _ := json.Marshal(ins)
	w.Write([]byte(responseString))
	return
}

type Args struct {
	UseCase ucincentives.UseCaseService
}

func New(args Args) *Controller {
	return &Controller{useCase: args.UseCase}
}
