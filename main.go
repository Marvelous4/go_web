package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", welcome)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal("ListenAndServer ::", err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go webserver codelab")
}
