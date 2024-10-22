package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Listening on port 8180...")
	http.ListenAndServe(":8180", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
