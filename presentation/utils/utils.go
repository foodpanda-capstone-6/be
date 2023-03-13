package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"vms-be/entities"
)

func GetRequestBody(r *http.Request) ([]byte, error) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return reqBody, nil
}
func GetUserLoginFieldsFromRequest(r *http.Request) (entities.LoginFields, error) {
	var fields entities.LoginFields
	r.Context()
	reqBody, err := GetRequestBody(r)
	if err != nil {
		return entities.LoginFields{}, err
	}

	err = json.Unmarshal(reqBody, &fields)
	if err != nil {
		return entities.LoginFields{}, err
	}
	return fields, err
}

func GetTokenFromRequest(r *http.Request) (string, error) {
	log.Println("[GetTokenFromRequest] not implemented")
	return "", nil
}
