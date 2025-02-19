package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	log.Println("Listening on port 8180...")
	http.ListenAndServe(":8180", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
