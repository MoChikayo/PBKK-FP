package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	httpDelivery "github.com/MoChikayo/PBKK-FP/pkg/delivery/http"
	"github.com/MoChikayo/PBKK-FP/pkg/domain"
	"github.com/MoChikayo/PBKK-FP/pkg/repository"
	"github.com/MoChikayo/PBKK-FP/pkg/routes"
	"github.com/MoChikayo/PBKK-FP/pkg/service"
)

func loadTemplates() *template.Template {
	tmpl := template.New("")

	// Load main templates
	mainPages, _ := filepath.Glob("templates/*.html")
	tmpl = template.Must(tmpl.ParseFiles(mainPages...))

	// Load components
	components, _ := filepath.Glob("templates/components/*.html")
	tmpl = template.Must(tmpl.ParseFiles(components...))

	return tmpl
}

func main() {
	os.Setenv("CGO_ENABLED", "0")
	// 1. Init DB

	db := config.InitDB()
	// 2. Auto migrate
	db.AutoMigrate(
		&domain.Book{},
		&domain.Customer{},
		&domain.Transaction{},
	)

	// 3. Init repositories
	bookRepo := repository.NewBookRepository()
	customerRepo := repository.NewCustomerRepository()
	transactionRepo := repository.NewTransactionRepository()

	// 4. Init services
	bookService := service.NewBookService(bookRepo)
	customerService := service.NewCustomerService(customerRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	// 5. Init handlers
	bookHandler := httpDelivery.NewBookHandler(bookService)
	customerHandler := httpDelivery.NewCustomerHandler(customerService)
	transactionHandler := httpDelivery.NewTransactionHandler(transactionService)

	// 6. Init Gin router
	r := gin.Default()

	r.SetHTMLTemplate(loadTemplates())

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})

	r.GET("/books", func(c *gin.Context) {
		c.HTML(200, "books.list.html", nil)
	})

	r.GET("/books/new", func(c *gin.Context) {
		c.HTML(200, "books.create.html", nil)
	})

	r.GET("/books/update/:id", func(c *gin.Context) {
		c.HTML(200, "books.update.html", gin.H{"ID": c.Param("id")})
	})

	r.GET("/customers", func(c *gin.Context) {
		c.HTML(200, "customers.list.html", nil)
	})

	r.GET("/transactions", func(c *gin.Context) {
		c.HTML(200, "transactions.list.html", nil)
	})
	// 7. Register all routes
	routes.RegisterRoutes(r, bookHandler, customerHandler, transactionHandler)

	// 8. Start server
	log.Println("Server running on http://localhost:9010")
	if err := r.Run(":9010"); err != nil {
		log.Fatal(err)
	}
}
