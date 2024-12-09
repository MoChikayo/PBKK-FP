package main

import (
	"log"
	"net/http"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Println("Server is running on http://localhost:9010")
<<<<<<< Updated upstream
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
=======
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
>>>>>>> Stashed changes
