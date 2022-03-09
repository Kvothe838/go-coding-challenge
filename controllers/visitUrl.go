package controllers

import (
	"codingchallenge/model"
	"log"
	"net/http"
	"net/url"
)

var visits []model.VisitedUrl

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

	visitedUrl := model.VisitedUrl{
		VisitorId: data.VisitorId,
		Url:       data.Url,
	}

	visits = append(visits, visitedUrl)
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
