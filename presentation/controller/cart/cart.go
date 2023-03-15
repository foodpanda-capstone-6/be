package market

import (
	"encoding/json"
	"log"
	"net/http"
	uccart "vms-be/usecase/cart"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	useCase uccart.UseCaseService
}

func (c *Controller) Routes(r chi.Router) {
	r.Get("/", c.getAll)
	r.Post("/purchase", c.purchase)
	r.Post("/add", c.update)
}

func (c *Controller) getAll(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerCart.getAll]")

	username := r.URL.Query().Get("username")

	cart := c.useCase.GetCart(username)

	responseString, err := json.Marshal(cart)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	log.Printf("getAll Cart %s \n", responseString)

	w.Write(responseString)
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) purchase(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerCart.purchase]")

	username := r.URL.Query().Get("username")

	c.useCase.Purchase(username)

	responseString := []byte("")
	w.Write(responseString)
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerCart.update]")

	/*
	   username: username,
	   id: openPurchaseDrawer.id,
	   quantity: voucherQty
	*/

	type Payload struct {
		Id       int `json:"id"`
		Quantity int
		Username string
	}

	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("[update cart controller] fail %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Printf("[update cart controller]  %v\n", payload)
	c.useCase.AddToCart(payload.Username, payload.Quantity, payload.Id)
	w.WriteHeader(http.StatusOK)
	return
}

type Args struct {
	UseCase uccart.UseCaseService
}

func New(args Args) *Controller {
	return &Controller{useCase: args.UseCase}
}
