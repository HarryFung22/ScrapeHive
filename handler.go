package main

import "net/http"

func handleRes(w http.ResponseWriter, r *http.Request) {
	resJson(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	resErr(w, 400, "Something went wrong")
}
