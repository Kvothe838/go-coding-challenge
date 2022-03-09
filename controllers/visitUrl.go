package controllers

import (
	"codingchallenge/services"
	"log"
	"net/http"
	"net/url"
)

func VisitUrl(response http.ResponseWriter, request *http.Request) {
	var data struct {
		VisitorId string `json:"visitorId"`
		Url       string `json:"url"`
	}

	err := Decode(request.Body, &data)
	if err != nil {
		log.Printf("error reading request body: '%v'", err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.VisitUrl(data.VisitorId, data.Url)

	if err != nil {
		log.Printf("error visiting url (visitorId = %v, url = %v): %v", data.VisitorId, data.Url, err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	Json(response, http.StatusOK, nil)
}

func GetVisitors(response http.ResponseWriter, request *http.Request) {
	urlParam := request.FormValue("url")

	if len(urlParam) == 0 {
		log.Printf("empty url read")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := url.QueryUnescape(urlParam)
	if err != nil {
		log.Printf("error uncoding url (urlParam = %v): %v", urlParam, err)
		return
	}

	totalVisitors, err := services.GetVisitors(url)
	if err != nil {
		log.Printf("error getting total visitors (url = %v): %v", url, err)
		return
	}

	responseData := struct {
		TotalVisitors int
	}{
		TotalVisitors: totalVisitors,
	}

	Json(response, http.StatusOK, responseData)
}
