package main

import (
	"github.com/cihanerman/SimpleMap/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routes.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
