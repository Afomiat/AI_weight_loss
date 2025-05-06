package main

import (
	"log"
	"time"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Afomiat/AI_weight_loss/config"
	"github.com/Afomiat/AI_weight_loss/delivery/router"
)

func main() {
    env := config.LoadEnv() // Load environment variables

    dbClient, err := config.ConnectDatabase(env) // Connect to MongoDB
    if err != nil {
        log.Fatalf("Could not connect to database: %v\n", err)
    }

    db := dbClient.Database(env.DBName) // Get the database instance

    // Create Gin router
    r := gin.Default()
    
    // Configure CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // You can specify specific origins here like []string{"http://localhost:3000"}
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    
    // Alternative simpler CORS configuration (less secure)
    // r.Use(cors.Default())

    // Setup your routes
    router.SetupRouter(r, env, db) // Assuming you modify SetupRouter to use the existing router

    if err := r.Run(":" + env.ServerPort); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    }
}