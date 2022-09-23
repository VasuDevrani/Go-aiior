package main

import (
	"fmt"
	"log"
	"mongo/routes"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server started...")

	fmt.Println("Listening at post 4000")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
