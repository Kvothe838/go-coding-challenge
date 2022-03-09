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

type setHandlerFunc func(path string, f http.HandlerFunc)

func buildSetHandleFunc(router *mux.Router, methods ...string) setHandlerFunc {
	return func(path string, f http.HandlerFunc) {
		router.HandleFunc(path, f).Methods(methods...)
	}
}

func handleRequests() {
	router := mux.NewRouter()

	Post := buildSetHandleFunc(router, "POST")
	Get := buildSetHandleFunc(router, "GET")

	Get("/getVisitors", controllers.GetVisitors)
	Post("/visit", controllers.VisitUrl)

	log.Println("serving in port 10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
