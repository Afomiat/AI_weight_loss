package config

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() (*mongo.Client) {
    client, err := mongo.NewClient(options.Client().ApplyURI(GetEnv("MONGO_URI")))
    if err != nil {
        log.Fatalf("Error creating MongoDB client: %v\n", err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v\n", err)
    }

    DB = client.Database(GetEnv("Ai_weight_loss"))

    log.Println("Connected to MongoDB!")
    return client
}
