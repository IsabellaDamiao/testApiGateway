package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":4040", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, "Hello World!")
}
