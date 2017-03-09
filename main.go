package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server is up on 8080 port")
	router := NewRouter()
	log.Fatalln(http.ListenAndServe(":8080", router))
}
