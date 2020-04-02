package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	print("Start Server")

	r := mux.NewRouter()

	r.HandleFunc("/hello", hello)
	r.HandleFunc("/world", world)

	http.ListenAndServe(":80", r)
}
