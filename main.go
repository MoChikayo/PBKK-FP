package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Dynamically load HTML templates
	var htmlFiles []string
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		// Check if the file is an HTML file
		if filepath.Ext(path) == ".html" {
			htmlFiles = append(htmlFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	if len(htmlFiles) == 0 {
		log.Fatalf("No HTML templates found in the 'templates' directory")
	}

	r.LoadHTMLFiles(htmlFiles...)

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	// Register Reset Database endpoint
	r.GET("/reset-database", ResetDatabaseEndpoint)

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
