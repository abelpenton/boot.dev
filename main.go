package main

import (
	"awesomeProject/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.FtpRouter()

	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
