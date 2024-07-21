package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func resJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to jsonify response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func resErr(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Intercepted a 500 error code: ", message)
	}
	type errRes struct {
		Error string `json:"error"`
	}

	resJson(w, code, errRes{
		Error: message,
	})
}
