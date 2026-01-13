package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello net/html")
}

func main() {
	http.HandleFunc("/", helloGoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
