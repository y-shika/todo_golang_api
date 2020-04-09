package handler

import "net/http"

// TODO: この部分を変えると恐らく登録していないパスにアクセスしたときの挙動が変わる。確認してみる
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not found", http.StatusNotFound)
}
