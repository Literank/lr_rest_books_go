package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a health endpoint handler
	router.GET("/", func(c *gin.Context) {
		// Return a simple response indicating the server is healthy
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
