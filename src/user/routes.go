package user

import "github.com/gin-gonic/gin"

func RegisterRoute(r *gin.Engine, handler *UserHandler) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", handler.RegisterUser)
	}
}
