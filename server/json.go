package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	// Print the error for the console if it is not nil
	if err != nil {
		log.Println(err)
	}

	//  Log all 5XX (server error) responses for easier debugging and visibility of server-side issues (e.g., 500 Internal Server Error).
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	// Make the struct with the encountered error, attact json elements so we can send that to the user
	type errorResponse struct {
		Error string `json:"error"`
	}

	// Call the respondWithJSON function to return the error struct to the user with the http code
	respondWithJSON(w, code, errorResponse{
	Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Set the response header to what type of content that is being returned. In this case json
	w.Header().Set("Content-Type", "application/json")

	// Marshal the data (package it for the client) and check for erros. If an error is found log it and return
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}

	// Set the http code and send the response
	w.WriteHeader(code)
	w.Write(dat)
}