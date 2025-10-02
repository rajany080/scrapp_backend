package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// createUserHandler handles the creation of a new user
func createUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
