package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajany080/scrapp_backend/models"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID      int    `json:"id" example:"1"`
	Name    string `json:"name" example:"Rajan"`
	Message string `json:"message" example:"User created successfully"`
}

// CreateUserHandler godoc
// @Summary      Create a new user
// @Description  Creates a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User data"
// @Success      200  {object}  UserResponse
// @Failure      400  {object}  map[string]string
// @Router       /api/users/signup [post]
func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Simple validation
		if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
			return
		}

		response := UserResponse{
			ID:      1, // Example static ID
			Name:    user.FirstName + " " + user.LastName,
			Message: "User created successfully",
		}
		c.JSON(http.StatusOK, response)
	}
}
