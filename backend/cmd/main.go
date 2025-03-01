package main

import (
    "log"
    "github.com/Afomiat/AI_weight_loss/backend/config"
    "github.com/Afomiat/AI_weight_loss/backend/delivery/router"
)

func main() {
    config.LoadEnv()
    config.ConnectDatabase()

    r := router.SetupRouter()
    if err := r.Run(":" + config.GetEnv("SERVER_PORT")); err != nil {
        log.Fatalf("Could not start server: %v\n", err)
    }
}
