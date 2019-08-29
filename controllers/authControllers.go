package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GoAuth/utils"

	"github.com/GoAuth/models"
)

var Register = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res := account.Create()
	utils.Response(w, res)
}

var Login = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res := models.Login(account.Email, account.Password)
	utils.Response(w, res)
}
