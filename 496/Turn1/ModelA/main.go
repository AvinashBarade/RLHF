package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a router
	router := gin.Default()

	// Define a GET endpoint
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
