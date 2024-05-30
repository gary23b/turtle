package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./cmd/runinweb")))
	fmt.Println("Golang: Serving directory")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// http://localhost:3000/
