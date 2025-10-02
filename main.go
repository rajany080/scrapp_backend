package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/rajany080/scrapp_backend/docs" // docs folder import for swagger

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Scrapp Backend API
// @version     1.0
// @description This is the backend API for Scrapp application built with Gin
// @host        localhost:8090
// @BasePath    /

// PingResponse represents the response for ping endpoint
type PingResponse struct {
	Message string `json:"message"`
}

// UserScrappResponse represents the response for user scrapp endpoint
type UserScrappResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

// Ping godoc
// @Summary      Health check endpoint
// @Description  Returns pong to verify the API is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  PingResponse
// @Router       /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(200, PingResponse{
		Message: "pong",
	})
}

// GetUserScrapp godoc
// @Summary      Get user scrapp data
// @Description  Retrieves scrapp data for a specific user
// @Tags         scrapp
// @Produce      json
// @Param        userId   path      string  true  "User ID"
// @Success      200      {object}  UserScrappResponse
// @Router       /scrapp/{userId} [get]
func getUserScrappHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")
	ctx.JSON(200, UserScrappResponse{
		Message: "Successfully hit the message",
		UserID:  userId,
	})
}

func main() {
	// Load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	router.GET("/ping", pingHandler)
	router.GET("/scrapp/:userId", getUserScrappHandler)

	// Start server on the configured port
	serverAddr := ":" + port
	fmt.Println("Server running on port: ", port)
	fmt.Println("Swagger docs available at: http://localhost:" + port + "/swagger/index.html")
	router.Run(serverAddr)
}
