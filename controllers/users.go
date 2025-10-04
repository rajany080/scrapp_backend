package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajany080/scrapp_backend/models"
	"github.com/rajany080/scrapp_backend/schemas"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUserHandler godoc
// @Summary      Create a new user
// @Description  Creates a new user in the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body schemas.CreateUserSchema true "User data"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/users/signup [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var create_user_body schemas.CreateUserSchema
		// Check if the provided body matches the schema and validation rules
		if err := c.ShouldBindJSON(&create_user_body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate custom rules (e.g., role validation)
		if err := create_user_body.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(create_user_body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Save the user to the database
		user := models.User{
			FirstName: create_user_body.FirstName,
			LastName:  create_user_body.LastName,
			Email:     create_user_body.Email,
			Phone:     create_user_body.Phone,
			Password:  string(hashedPassword),
			About:     create_user_body.About,
			Role:      create_user_body.Role,
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		response := gin.H{
			"message": "User created successfully",
		}
		c.JSON(http.StatusOK, response)
	}
}

// Login godoc
// @Summary      User login
// @Description  Authenticates a user with email and password
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        credentials body schemas.LoginSchema true "Login credentials"
// @Success      200  {object}  schemas.LoginResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /api/users/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginBody schemas.LoginSchema

		// Validate request body
		if err := c.ShouldBindJSON(&loginBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find user by email
		var user models.User
		if err := db.Where("email = ?", loginBody.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Compare password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Return user data (without password)
		response := schemas.LoginResponse{
			Message: "Login successful",
			User: schemas.GetUserResponse{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Phone:     user.Phone,
				About:     user.About,
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetUsersHandler godoc
// @Summary      Get users
// @Description  Fetches users from the system with pagination
// @Tags         Users
// @Produce      json
// @Param        page     query     int  false  "Page number"  default(1)
// @Param        pageSize query     int  false  "Items per page" default(50)
// @Success      200  {array}   schemas.GetUserResponse
// @Failure      500  {object}  map[string]string
// @Router       /api/users [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse pagination query parameters
		page := 1
		pageSize := 50
		if p := c.Query("page"); p != "" {
			fmt.Sscanf(p, "%d", &page)
		}
		if ps := c.Query("pageSize"); ps != "" {
			fmt.Sscanf(ps, "%d", &pageSize)
		}
		offset := (page - 1) * pageSize

		// Fetch only required fields
		var users []models.User
		if err := db.
			Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		// Map to DTOs
		var response []schemas.GetUserResponse
		for _, u := range users {
			response = append(response, schemas.GetUserResponse{
				ID:        u.ID,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
				Phone:     u.Phone,
				About:     u.About,
				Role:      u.Role,
				CreatedAt: u.CreatedAt,
				UpdatedAt: u.UpdatedAt,
			})
		}

		// Return paginated response
		c.JSON(http.StatusOK, gin.H{
			"page":     page,
			"pageSize": pageSize,
			"users":    response,
		})
	}
}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Fetches a single user by their ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  schemas.GetUserResponse
// @Failure      404  {object}  map[string]string
// @Router       /api/users/{id} [get]
func GetUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Convert to response DTO
		response := schemas.GetUserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			About:     user.About,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		c.JSON(http.StatusOK, response)
	}
}
