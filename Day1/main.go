package main

import (
	"log"
	"net/http"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello Web")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", ViewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
