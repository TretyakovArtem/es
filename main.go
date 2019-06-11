package main

import (
	"log"
	"net/http"

	"github.com/es/controllers"
	"github.com/es/driver"

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

	router.HandleFunc("/", controller.GenerateSeq(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
