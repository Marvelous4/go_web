package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", welcome)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go webserver codelab")
}
