package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GoAuth/app"
	"github.com/GoAuth/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.Register).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/contact/new", controllers.AddContact).Methods("POST")
	router.HandleFunc("/api/contact/show", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication) //melampirkan jwt middleware

	port := os.Getenv("port")
	if port == "" {
		port = "8000"
	}

	log.Println(port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
