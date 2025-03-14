package main

import (
	"go-project-management-book/infra"
	"go-project-management-book/src/book"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	infra.ConnectionDB()

	repo := book.NewBookRepository(infra.DB)
	service := book.NewBookService(repo)
	handler := book.NewBookHandler(service)

	book.RegisterRoute(r, handler)

	r.Run(":8080")
}
