package routes

import (
	httpDelivery "github.com/MoChikayo/PBKK-FP/pkg/delivery/http"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	bh *httpDelivery.BookHandler,
	ch *httpDelivery.CustomerHandler,
	th *httpDelivery.TransactionHandler,
) {
	api := r.Group("/api")

	books := api.Group("/books")
	{
		books.GET("", bh.GetBooks)
		books.GET("/:id", bh.GetBook)
		books.POST("", bh.CreateBook)
		books.PUT("/:id", bh.UpdateBook)
		books.DELETE("/:id", bh.DeleteBook)
	}

	customers := api.Group("/customers")
	{
		customers.GET("", ch.GetCustomers)
		customers.GET("/:id", ch.GetCustomer)
		customers.POST("", ch.CreateCustomer)
		customers.PUT("/:id", ch.UpdateCustomer)
		customers.DELETE("/:id", ch.DeleteCustomer)
	}

	transactions := api.Group("/transactions")
	{
		transactions.GET("", th.GetTransactions)
		transactions.GET("/:id", th.GetTransaction)

		transactions.POST("/borrow", th.Borrow)
		transactions.PUT("/return/:id", th.Return)

		transactions.DELETE("/:id", th.DeleteTransaction)
	}

}
