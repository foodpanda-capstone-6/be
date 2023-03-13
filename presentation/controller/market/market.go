package market

import (
	"encoding/json"
	"log"
	"net/http"
	ucmarket "vms-be/usecase/market"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase ucmarket.UseCaseService
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/check", c.hello)
	r.Get("/", c.getVoucherTypeAmount)
	r.Post("/", c.login)

}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from market"))
	return
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerMarket.hello]")

	return
}

func (c *Controller) getVoucherTypeAmount(w http.ResponseWriter, r *http.Request) {
	c.getVoucherAll(w, r)
	return
}

func (c *Controller) getVoucherAll(w http.ResponseWriter, r *http.Request) {

	log.Println("[getVoucherAll]")

	mvs := c.useCase.GetMarketVouchers()

	responseString, _ := json.Marshal(mvs)

	w.Write([]byte(responseString))

	return
}

type Args struct {
	UseCase ucmarket.UseCaseService
}

func New(args Args) *Controller {
	return &Controller{useCase: args.UseCase}
}
