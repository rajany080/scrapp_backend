package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajany080/scrapp_backend/controllers"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	users := router.Group("/users")
	{
		users.POST("/login", controllers.Login(db))
		users.POST("/signup", controllers.CreateUser(db))
		users.GET("", controllers.GetUsers(db))
		users.GET("/:id", controllers.GetUserById(db))
	}
}
