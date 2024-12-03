package main

import (
	"log"
	"net/http"

	"github.com/MoChikayo/PBBK-FP/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Println("Server is running on http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
