package main

import (
	"log"
	"net/http"

	"github.com/jittakal/go-apps/ams/router"
)

func main() {
	router := router.NewRouter()
	log.Println("Starting server, access it with 'http://localhost:8080/{api-routes}'")
	log.Fatal(http.ListenAndServe(":8080", router))
}
