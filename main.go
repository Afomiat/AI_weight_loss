package main

import (
	"log"

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

    r := router.SetupRouter(env, db) // Pass `db` correctly
    if err := r.Run(":" + env.ServerPort); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    }
}

