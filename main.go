package main

import (
	"log"
	dependencyinjection "main/dependency-injection"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
