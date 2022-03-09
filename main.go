package main

import (
	"codingchallenge/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/visit", controllers.VisitUrl).Methods("POST")
	myRouter.HandleFunc("/getVisitors", controllers.GetVisitors).Methods("GET")

	log.Println("serving in port 10000")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
