package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receiving request for publicgo")
	fmt.Fprintf(w, "Hello from publicgo")
}

func pingInternalPage(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("Calling internalgo app using internal dns")
	fmt.Fprintf(w, "Here's the response from internalgo: %s", parsedBody)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pinginternal", pingInternalPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
