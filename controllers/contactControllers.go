package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GoAuth/models"
	"github.com/GoAuth/utils"
)

var AddContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)

	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utils.Response(w, utils.Message(false, "Error decoding request body"))
	}

	log.Println("user id = ", user)
	contact.UserId = user
	res := contact.Create()
	utils.Response(w, res)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetAllContacts(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Response(w, resp)
}
