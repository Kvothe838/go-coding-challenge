package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"codingchallenge/model"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/visit", visitUrl).Methods("POST")
	myRouter.HandleFunc("/getVisitors", getVisitors).Methods("GET")
	log.Println("serving in port 10000")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

var visits []model.VisitedUrl

func visitUrl(response http.ResponseWriter, request *http.Request) {
	var data struct {
		VisitorId string `json:"visitorId"`
		Url       string `json:"url"`
	}

	err := decode(request.Body, &data)
	if err != nil {
		log.Printf("error reading request body: '%v'", err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	visitedUrl := model.VisitedUrl{
		VisitorId: data.VisitorId,
		Url:       data.Url,
	}

	visits = append(visits, visitedUrl)
	Json(response, http.StatusOK, nil)
}

func getVisitors(response http.ResponseWriter, request *http.Request) {
	urlParam := request.FormValue("url")

	if len(urlParam) == 0 {
		log.Printf("empty url read")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := url.QueryUnescape(urlParam)
	if err != nil {
		log.Fatal(err)
		return
	}

	distinctVisitors := make(map[string]bool, 0)

	for _, visit := range visits {
		if visit.Url == url {
			distinctVisitors[visit.VisitorId] = true
		}
	}

	totalVisitors := 0

	for _, hasVisitedUrl := range distinctVisitors {
		if hasVisitedUrl {
			totalVisitors++
		}
	}

	responseData := struct {
		TotalVisitors int
	}{
		TotalVisitors: totalVisitors,
	}

	Json(response, http.StatusOK, responseData)
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
