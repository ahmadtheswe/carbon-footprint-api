package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadtheswe/carbon-footprint-api/config"
	"github.com/ahmadtheswe/carbon-footprint-api/router"
)

func main() {
	// Load configuration from config.yaml
	config.LoadConfig()

	// Initialize database connection
	db := config.ConnectDatabase()

	// Keep database connection alive (you can use this in handlers later)
	_ = db

	// Setup routes
	router.SetupRoutes()

	port := ":" + config.AppConfig.Port
	fmt.Printf("Server is listening on port %s...\n", config.AppConfig.Port)
	http.ListenAndServe(port, nil)
}
