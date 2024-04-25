package main

import (
	"log"

	"github.com/Ephrem-shimels21/GoCrudChallenge/middlewares"
	"github.com/Ephrem-shimels21/GoCrudChallenge/routes"
	"github.com/Ephrem-shimels21/GoCrudChallenge/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Apply CORS middleware to handle cross-origin requests
	router.Use(middlewares.CORS())

	// Apply middleware to handle internal server errors
	router.Use(middlewares.Handle500())

	// Initialize the in-memory storage
	personStorage := storage.NewInMemoryPersonStorage()

	// Set up the routes for the 'person' resource
	routes.SetupPersonRoutes(router, personStorage)

	// Handle non-existing routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Resource not found"})
	})

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
