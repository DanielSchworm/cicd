package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorldHandler() {
	http.HandleFunc("/", printHelloWorld)
}

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	HelloWorldHandler()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
