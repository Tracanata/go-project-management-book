package book

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, handler *BookHandler) {
	bookRoutes := r.Group("/book")
	{
		bookRoutes.GET("/", handler.GetAllBooks)
		bookRoutes.GET("/:id", handler.GetBookById)
		bookRoutes.POST("/add", handler.AddBook)
		bookRoutes.PUT("/edit", handler.UpdateBook)
		bookRoutes.DELETE("/delete/:id", handler.DeleteBook)
	}

}
