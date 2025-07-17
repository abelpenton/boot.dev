package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"boot.dev/internal/routes"
)

func main() {
	config := &routes.ApiConfig{
		Hits: atomic.Int32{},
	}

	router := routes.FtpRouter(config)

	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
