package routes

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	router.POST("/users", controllers.createUserHandler)
}
