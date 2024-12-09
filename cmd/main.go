package main

import (
	"log"

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
