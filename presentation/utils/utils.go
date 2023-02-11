package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"vms-be/entities"
)

func GetUserLoginFieldsFromRequest(r *http.Request) (entities.LoginFields, error) {
	var fields entities.LoginFields

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return entities.LoginFields{}, err
	}
	err = json.Unmarshal(reqBody, &fields)
	if err != nil {
		return entities.LoginFields{}, err
	}
	return fields, err
}
