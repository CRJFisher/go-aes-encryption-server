package main

import (
	"log"
	"net/http"
)

func main() {

	router := MuxRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
