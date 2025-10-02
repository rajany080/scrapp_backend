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

	"github.com/rajany080/scrapp_backend/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	dsn := "host=localhost user=postgres password=test dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	fmt.Println("Database connected!", db)

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	apiGroup := router.Group("/api")
	routes.UserRoutes(apiGroup)

	// Start server on the configured port
	serverAddr := ":" + port
	fmt.Println("Server running on port: ", port)
	fmt.Println("Swagger docs available at: http://localhost:" + port + "/swagger/index.html")
	router.Run(serverAddr)
}
