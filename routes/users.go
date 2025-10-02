package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajany080/scrapp_backend/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/users/signup", controllers.CreateUserHandler)
	router.POST("/users/login", func(ctx *gin.Context) {
		// Login logic here
		ctx.JSON(200, gin.H{"message": "Login successful"})
	})
	router.GET("/users", controllers.CreateUserHandler)
	router.GET("/users/:id", func(ctx *gin.Context) {
		// Get user by ID logic here
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{"message": "User fetched successfully", "id": id})
	})
}
