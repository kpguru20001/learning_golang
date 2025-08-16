package main

import (
	"log"
	"net/http"
	"os"
	"shadow-docs/configs"
	"shadow-docs/internal/database"
	"shadow-docs/internal/handlers"
	"shadow-docs/pkg/middleware"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadConfig()
	if err := database.Connect(); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		os.Exit(1)
	}
	defer database.Disconnect()

	clerk.SetKey(configs.Configuration.Clerk.SecretKey)
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.Use(middleware.GinContextToContextMiddleware())

	// Define routes
	r.GET("/", homeHandler)
	r.GET("/health", healthHandler)
	r.GET("/graphql", handlers.GQLPlaygroundHandler())
	r.POST("/graphql", handlers.GraphQLHandler())

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Shadow Docs API",
	})
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
