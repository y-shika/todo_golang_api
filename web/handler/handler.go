package handler

import "net/http"

// NotFoundHandler is a default handler for the undefined resource path.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not found", http.StatusNotFound)
}
