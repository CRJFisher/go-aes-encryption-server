package main

import (
	"log"
	"net/http"
)

func main() {

	StartDB()
	defer StopDB()

	router := MuxRouter()

	log.Fatal(http.ListenAndServe(":8082", router))

}
