package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		name, _ := os.Hostname()
		log.Println(name, "ping")
		fmt.Fprint(writer, name+" pong")
	})
	http.HandleFunc("/service", func(writer http.ResponseWriter, request *http.Request) {
		// 对应的服务
		resp, _ := http.Get("http://k8s-combat-service:8080/ping")
		body, _ := io.ReadAll(resp.Body)
		log.Println(string(body))
		fmt.Fprint(writer, string(body))
	})
	http.ListenAndServe(":8080", nil)
}
