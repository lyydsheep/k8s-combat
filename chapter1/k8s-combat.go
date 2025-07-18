package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("pong")
		fmt.Fprint(writer, "pong")
	})
	http.ListenAndServe(":8080", nil)
}
