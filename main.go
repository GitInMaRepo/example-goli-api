package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8080", nil)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	origin := request.Header.Get("Origin")
	writer.Header().Set("Access-Control-Allow-Origin", origin)
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	message := "Welcome to GO|i-API Version 0.1"
	body, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	writer.Write(body)
}
