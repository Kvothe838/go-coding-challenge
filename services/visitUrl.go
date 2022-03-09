package services

import (
	"codingchallenge/model"
	"sync"
)

var visits []model.VisitedUrl
var visitsMutex = sync.Mutex{}

func VisitUrl(visitorId string, url string) error {
	visitedUrl := model.VisitedUrl{
		VisitorId: visitorId,
		Url:       url,
	}

	visitsMutex.Lock()

	visits = append(visits, visitedUrl)

	visitsMutex.Unlock()

	return nil
}

func GetVisitors(url string) (int, error) {
	distinctVisitors := make(map[string]bool, 0)

	visitsMutex.Lock()

	// Set VisitorId as map key in true when visited url matches url param to avoid repeating visits for same visitor id and visited url
	for _, visit := range visits {
		if visit.Url == url {
			distinctVisitors[visit.VisitorId] = true
		}
	}

	visitsMutex.Unlock()

	totalVisitors := 0

	// Every entry of the map is a different visitor for the url param only if value is true
	for _, hasVisitedUrl := range distinctVisitors {
		if hasVisitedUrl {
			totalVisitors++
		}
	}

	return totalVisitors, nil
}
