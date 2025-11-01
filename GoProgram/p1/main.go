package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	port := ":8081"
	fmt.Printf("Starting server on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,World!")
}
