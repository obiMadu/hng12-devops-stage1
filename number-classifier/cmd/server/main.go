package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/number-classifier/internal/api"
	"github.com/number-classifier/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	
	// Add middleware
	router.Use(middleware.CORS())
	
	// Routes
	router.GET("/health", api.HealthCheck)
	router.GET("/api/classify-number", api.ClassifyNumber)

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
