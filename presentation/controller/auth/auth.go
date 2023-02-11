package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"vms-be/presentation/utils"

	uc_auth "vms-be/usecase/auth"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	usecase uc_auth.UseCaseInterface
}

type Args struct {
	UseCase uc_auth.UseCaseInterface
}

func (c *Controller) Routes(r chi.Router) {

	r.Get("/check", c.hello)
	r.Post("/", c.login)

}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from auth"))
	return
}

var ErrInvalidTokenInput error = errors.New("invalid token")

func GetToken(r *http.Request) (*string, error) {
	ca, err := r.Cookie("Authorization")
	if err != nil {
		return nil, ErrInvalidTokenInput
	}
	return &ca.Value, nil
}

type LoginResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Data struct {
	Token string `json:"token"`
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	log.Println("[ControllerAuth.hello]")

	loginFields, err := utils.GetUserLoginFieldsFromRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		responseBody := fmt.Sprintf("{ \"message\": \"%v\" }", err.Error())
		w.Write([]byte(responseBody))
		return
	}
	jwt_String, err := c.usecase.Login(loginFields)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseBody_JSON := LoginResponse{Message: err.Error()}
		responseBody, _ := json.Marshal(responseBody_JSON)

		w.Write([]byte(responseBody))
		return
	}

	responseBody_JSON := LoginResponse{Data: Data{Token: jwt_String.String()}}

	responseBody, err := json.Marshal(responseBody_JSON)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseBody := fmt.Sprintf("{ \"message\": \"%v\" }", err.Error())
		w.Write([]byte(responseBody))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
	return
}

func New(args Args) *Controller {
	return &Controller{usecase: args.UseCase}
}
