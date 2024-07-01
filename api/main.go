package main

import (
	"api/config"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Init()

	router := router.Init()

	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
