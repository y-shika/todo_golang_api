package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type fakeGithubData struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatarUrl"`
}

func fakeGithubAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fakeGithubData{
			Login:     "y-shika",
			ID:        1,
			AvatarURL: "https://avatars3.githubusercontent.com/u/583231?v=4",
		}

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	}
}

func helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	}
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Printf("Starting server at Port %s", port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})

	r.Handle("/api/github", fakeGithubAPIHandler()).Methods(http.MethodGet)
	r.Handle("/hello", helloHandler()).Methods(http.MethodGet)

	r.HandleFunc("/world", world)

	http.ListenAndServe(":"+port, r)
}
