package main

import (
	"go-project-management-book/infra"
	"go-project-management-book/src/book"
	"go-project-management-book/src/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	infra.ConnectionDB()

	// Book
	repoBook := book.NewBookRepository(infra.DB)
	serviceBook := book.NewBookService(repoBook)
	handlerBook := book.NewBookHandler(serviceBook)

	book.RegisterRoute(r, handlerBook)

	// User
	repoUser := user.NewUserRepository(infra.DB)
	serviceUser := user.NewUserService(repoUser)
	handlerUser := user.NewUserHandler(serviceUser)

	user.RegisterRoute(r, handlerUser)
	r.Use(gin.Recovery())

	r.Run(":8080")
}
