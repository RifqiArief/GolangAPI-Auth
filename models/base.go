package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	//loading .env file
	err := godotenv.Load("config/database.env")
	if err != nil {
		log.Println("models/base/line:19")
		log.Fatal(err)
	}

	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	log.Println(dbUri)

	//open connection
	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		log.Println("models/base/line:34")
		log.Fatal(err)
	}

	db = conn
	db.LogMode(true)
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
