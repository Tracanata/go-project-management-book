package book

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, handler *BookHandler) {
	bookRoutes := r.Group("/book")
	{
		bookRoutes.GET("/", handler.GetAllBooks)
		bookRoutes.GET("/:codeBook", handler.GetBookByCodeBook)
		bookRoutes.POST("/add", handler.AddBook)
		bookRoutes.PUT("/edit/:codeBook", handler.UpdateBook)
		bookRoutes.DELETE("/delete/:codeBook", handler.DeleteBook)
	}

}
