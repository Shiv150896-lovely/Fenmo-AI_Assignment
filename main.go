package main

import (
	"fenmo-ai-assignment/config"
	"fenmo-ai-assignment/database"
	"fenmo-ai-assignment/routes"
	"fmt"
	"log"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()

	// Initialize database
	if err := database.Init(cfg.DBPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Setup routes
	router := routes.SetupRoutes()

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
