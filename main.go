package main

import (
	"halyklife-lib/config"
	"halyklife-lib/routes"
	"log"
	"net/http"
)

func main() {
	err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := routes.SetupRoutes()
	log.Printf("Starting the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
