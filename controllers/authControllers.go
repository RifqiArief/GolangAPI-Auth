package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GoAuth/utils"

	"github.com/GoAuth/models"
)

var Register = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	account := &models.Account{}
	err := json.Unmarshal(body, account)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res := account.Create()
	resLog, _ := json.Marshal(res)
	utils.Logging.Println(resLog)
	utils.Response(w, res)
}

var Login = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	account := &models.Account{}

	err := json.Unmarshal(body, account)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res := models.Login(account.Email, account.Password)
	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}
