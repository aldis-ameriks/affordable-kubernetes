package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	fmt.Fprintf(w, "Welcome to Kubernetes!")
}

func internalPage(w http.ResponseWriter, r *http.Request) {
	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := netClient.Get("http://internalgo.default.svc.cluster.local:8081")
	if err != nil {
		fmt.Println("Error", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	var parsedBody = string(body)
	fmt.Println(parsedBody)
	fmt.Println("Internal request received")
	fmt.Fprintf(w, "Welcome to page that calls an internal service!! %s", parsedBody)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/internal", internalPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
