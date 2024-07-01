package main

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := router.Init()

	fmt.Println("Listening on port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
