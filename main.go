package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/visit", visitUrl).Methods("POST")
	log.Println("serving in port 10000")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type VisitUrlData struct {
	VisitorId  string `json:"visitorId"`
	VisitedUrl string `json:"visitedUrl"`
}

var visits []VisitUrlData

func visitUrl(response http.ResponseWriter, request *http.Request) {
	var data VisitUrlData

	err := decode(request.Body, &data)
	if err != nil {
		log.Printf("error reading request body: '%v'", err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	visits = append(visits, data)
	Json(response, http.StatusOK, nil)
}

func decode(reader io.Reader, val interface{}) error {
	err := json.NewDecoder(reader).Decode(val)
	if err != nil {
		log.Printf("error decoding %T, error: %s", val, err.Error())
	}

	return err
}

func Json(response http.ResponseWriter, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("error while mashalling object %v: %+v", data, err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)

	_, err = response.Write(bytes)
	if err != nil {
		log.Printf("error while writting bytes to response writer: %+v", err)
	}
}
