package routes

import (
	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/controllers"
	"github.com/gorilla/mux"
)

<<<<<<< Updated upstream
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
=======
var RegisterBookStoreRoutes = func(router *gin.Engine) {
	// Book routes
	router.POST("/book", controllers.CreateBook)
	router.GET("/book", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)

	// customer routes
	// router.POST("/customer", func(c *gin.Context) {
	// 	controllers.Createcustomer(c)
	// })
	router.POST("/customer", controllers.CreateCustomer)
	router.GET("/customer", controllers.GetAllCustomers)
	router.GET("/customer/:customerId", controllers.GetCustomerById)
	router.PUT("/customer/:customerId", controllers.UpdateCustomer)
	router.DELETE("/customer/:customerId", controllers.DeleteCustomer)

	// Transaction routes
	// router.POST("/transaction", func(c *gin.Context) {
	// 	controllers.CreateTransaction(c)
	// })
	router.POST("/transaction", controllers.CreateTransaction)
	router.GET("/transaction", controllers.GetAllTransactions)
	router.GET("/transaction/:transactionId", controllers.GetTransactionById)
	router.PUT("/transaction/:transactionId/status", controllers.UpdateTransactionStatus)

	// Database reset route
	router.POST("/reset-database-endpoint", func(c *gin.Context) {
		if err := config.ResetDatabaseEndpoint(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "database reset successful"})
	})
>>>>>>> Stashed changes
}
