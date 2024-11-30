package routes

import (
	"github.com/MoChikayo/PBBK-FP/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookStoreRoutes(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}
