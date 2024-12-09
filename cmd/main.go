package main

import (
	"log"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	// Start the server
	log.Println("Server is running on http://localhost:9010")
	log.Fatal(r.Run(":9010"))
}

// ResetDatabaseEndpoint handles the database reset request
func ResetDatabaseEndpoint(c *gin.Context) {
	if err := config.ResetDatabaseEndpoint(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Database reset successfully"})
}
