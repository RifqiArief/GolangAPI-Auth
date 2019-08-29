package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoAuth/models"
	"github.com/GoAuth/utils"
)

var AddContact = func(w http.ResponseWriter, r *http.Request) {

	//logging request
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	utils.Logging.Println(string(body))

	user := r.Context().Value("user").(uint)

	contact := &models.Contact{}

	err := json.Unmarshal(body, contact)
	if err != nil {
		utils.Response(w, utils.Message(false, "Error decoding request body"))
	}

	log.Println("user id = ", user)
	contact.UserId = user
	res := contact.Create()

	//logging response
	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	utils.Logging.Println("GET all contact")
	id := r.Context().Value("user").(uint)
	data := models.GetAllContacts(id)
	res := utils.Message(true, "success")
	res["data"] = data
	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}
