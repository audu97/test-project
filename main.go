package main

import (
	"log"
	"net/http"
)

func firstEndPointHandler(w http.ResponseWriter, r *http.Request) {
	message := "this is the first endpoint"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func secondEndPointHandler(w http.ResponseWriter, r *http.Request) {
	message := "second endpoint"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/first", firstEndPointHandler)
	http.HandleFunc("/second", secondEndPointHandler)
	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
