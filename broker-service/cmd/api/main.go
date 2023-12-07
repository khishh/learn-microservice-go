package main

import (
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting API server on port at %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	// start the server

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
