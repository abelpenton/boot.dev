package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../../index.html"))
}
