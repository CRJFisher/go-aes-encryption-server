package main

import (
	"log"
	"net/http"
)

// UseRemoteStore toggles the microservice store, defaulting to in-memory store
var UseRemoteStore = true

func main() {

	router := MuxRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	// log.Fatal(http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", router))

}
