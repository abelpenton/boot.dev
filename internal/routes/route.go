package routes

import (
	"net/http"
)

func FtpRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app", http.FileServer(http.Dir("client"))).ServeHTTP(w, r)
	})

	router.HandleFunc("/app/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
	})

	router.HandleFunc("/", readinessHandler)

	return router
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
