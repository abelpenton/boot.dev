package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app", http.FileServer(http.Dir("."))).ServeHTTP(w, r)
	})

	router.HandleFunc("/app/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
	})

	router.HandleFunc("/", readinessHandler)

	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
