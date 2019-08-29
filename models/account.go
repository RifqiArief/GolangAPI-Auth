package models

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/GoAuth/data"
	"github.com/GoAuth/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Nama     string `json:"nama" sql:"type:varchar(30)"`
	Alamat   string `json:"alamat" sql:"type:varchar(255)"`
	Kota     string `json:"kota" sql:"type:varchar(30)"`
	NoTelp   string `json:"no_telp" sql:"type:varchar(20)"`
	Email    string `json:"email" sql:"type:varchar(100)"`
	Password string `json:"password" sql:"type:varchar(255)"`
	Image    string `json:"image" sql:"type:varchar(255)"`
	Status   bool   `json:"status"`
	Token    string `json:"token "sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email is required"), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	//menampung data dari database
	temp := &Account{}

	//cek error dan duplicate email
	err := GetDB().Table("accounts").Where("email=?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error, please retry"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email already use"), false
	}

	return utils.Message(false, "Required passed"), true
}

func (account *Account) Create() map[string]interface{} {

	res, ok := account.Validate()
	if !ok {
		return res
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Message(false, err.Error())
	}

	account.Password = string(hash)

	//query insert into account
	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	//create jwt token untuk yang baru teregistrasi
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return utils.Message(false, "Failed generate token string")
	}
	account.Token = tokenString

	data := &data.Register{
		Nama:  account.Nama,
		Email: account.Email,
		Token: account.Token,
	}
	response := utils.Message(true, "User has ben created")
	response["data"] = data
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}

	//get account dengan email

	err := GetDB().Table("accounts").Where("email =?", email).First(account).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return utils.Message(false, "Email not registered")
		}
		return utils.Message(false, "Connection error, please retry")
	}

	//chek password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Wrong password, please try again")
	}

	tk := &Token{}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return utils.Message(false, "Failed generate token string")
	}

	account.Token = tokenString

	data := &data.Login{
		Nama:   account.Nama,
		Alamat: account.Alamat,
		Kota:   account.Kota,
		NoTelp: account.NoTelp,
		Email:  account.Email,
		Image:  account.Image,
		Token:  account.Token,
	}

	response := utils.Message(true, "Login success")
	response["data"] = data
	return response
}
