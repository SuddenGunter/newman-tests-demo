package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status string `json:"status"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsBody, _ := json.Marshal(response{Status: "OK"})
	w.Write(jsBody)
}

func main() {
	http.HandleFunc("/sample", hello)
	http.ListenAndServe(":8080", nil)
}
