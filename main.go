package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println("Request received")
	fmt.Fprintf(w, "Welcome to Kubernetes! %s", parsedBody)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
