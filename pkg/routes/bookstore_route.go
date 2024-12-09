package routes

import (
	"github.com/MoChikayo/PBKK-FP/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	// Book routes
	router.POST("/book", controllers.CreateBook)
	router.GET("/book", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)

	// User routes
	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.GetAllUsers)
	router.GET("/user/:userId", controllers.GetUserById)

	// Transaction routes
	router.POST("/transaction", controllers.CreateTransaction)
	router.GET("/transaction", controllers.GetAllTransactions)
	router.GET("/transaction/:transactionId", controllers.GetTransactionById)
}
