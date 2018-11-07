package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	fmt.Fprintf(w, "Welcome to Kubernetes!")
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
