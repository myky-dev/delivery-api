package main

import (
	"log"
	"net/http"

	"github.com/myky-dev/delivery-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("server starting on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}