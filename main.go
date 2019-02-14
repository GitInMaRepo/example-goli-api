package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(getPort(), nil)
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

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
