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

}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from incentives"))
	return
}

func (c *Controller) getVoucherTypeAmount(w http.ResponseWriter, r *http.Request) {
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
