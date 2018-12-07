package main

import (
	"log"
	"net/http"

	"github.com/Howlyao/Server/service"
)

func main() {
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
