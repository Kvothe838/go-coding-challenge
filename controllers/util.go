package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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

func Decode(reader io.Reader, val interface{}) error {
	err := json.NewDecoder(reader).Decode(val)
	if err != nil {
		log.Printf("error decoding %T, error: %s", val, err.Error())
	}

	return err
}
