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

}
