package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	print("Start Server")

	r := mux.NewRouter()

	r.Handle("/api/github", fakeGithubAPIHandler()).Methods(http.MethodGet)
	r.Handle("/hello", helloHandler()).Methods(http.MethodGet)

	r.HandleFunc("/world", world)

	http.ListenAndServe(":8080", r)
}
