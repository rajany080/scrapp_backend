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
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	sslMode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, sslMode)
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
