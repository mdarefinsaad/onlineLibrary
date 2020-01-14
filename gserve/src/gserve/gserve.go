package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Gserve Started")
	// Start the server
	log.Fatal(http.ListenAndServe(":7070", nil))
}
