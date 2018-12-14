package main

import (
	"log"
	"net/http"

	"github.com/Howlyao/REST_API_Server/service"
)

func main() {
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
