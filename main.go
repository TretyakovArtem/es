package main

import (
	"esequence/controllers"
	"esequence/driver"
	"esequence/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	db := driver.ConnectDB()

	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/", controller.GenerateSeq(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
