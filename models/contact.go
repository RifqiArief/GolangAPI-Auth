package models

import (
	"log"

	"github.com/GoAuth/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Nama   string `json:"nama"`
	NoTelp string `json:"no_telp"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {

	if contact.Nama == "" {
		return utils.Message(false, "Nama harus di isi"), false
	}

	if contact.NoTelp == "" {
		return utils.Message(false, "Nomer telepon belum di isi"), false
	}

	if contact.UserId <= 0 {
		return utils.Message(false, "Kontak tidak di ketahui"), true
	}

	return utils.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	res, ok := contact.Validate()
	if !ok {
		return res
	}

	GetDB().Create(contact)

	response := utils.Message(true, "success")
	response["contact"] = contact
	return response
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id=?", id).First(contact).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return contact
}

func GetAllContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Find(&contacts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return contacts
}
