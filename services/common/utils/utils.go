package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	log.Println("error : ", err)
	WriteJson(w, status, map[string]string{
		"error": err.Error(),
	})
}

func ParseJson(req *http.Request, v any) error {
	if req.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(req.Body).Decode(v)
}
