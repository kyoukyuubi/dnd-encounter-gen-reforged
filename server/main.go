package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	platform string
}

func main() {
	// set the port
	const port = "8080"
	
	// get the .env file
	godotenv.Load()

	// get the platform from the .env file and check if it is empty
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set!")
	}

	// store the data from the .env to a config struct
	cfg := apiConfig {
		platform: platform,
	}

	// setup the MuxServer for the api
	mux := http.NewServeMux()

	// start the server
	server := http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	
	log.Printf("API is running on port %s. Platform is set to %s", port, cfg.platform)
	log.Fatal(server.ListenAndServe())
}