package main

import (
	"daveslist-emdpcv/api/database"
	"daveslist-emdpcv/api/middlewares"
	"daveslist-emdpcv/api/routes"
	"daveslist-emdpcv/api/settings"
	"fmt"
	"log"
	"time"

	"daveslist-emdpcv/api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Daveslist API
// @version 1.0
// @description This is a sample server for a used car listing application.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	config := settings.LoadEnv()

	// Initialize the database
	database.InitDatabase()

	// Set up the Gin router
	r := gin.Default()

	// Set up CORS middleware options
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup error logger middleware
	r.Use(middlewares.ErrorLogger())

	// Setup routes
	routes.Setup(r)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Daveslist API"
	docs.SwaggerInfo.Description = "This is a sample server for a used car listing application."
	docs.SwaggerInfo.Version = config.ApiVersion
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = config.ApiPath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Swagger endpoint
	log.Printf("Swagger endpoint: http://%s:%s/swagger/index.html", config.Host, config.Port)
	swaggerURL := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", config.Host, config.Port))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// Run the server
	log.Printf("Server is running on %s:%s", config.Host, config.Port)
	log.Fatal(r.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
