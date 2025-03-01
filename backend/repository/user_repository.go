package repository

import (
    "context"
    "github.com/Afomiat/AI_weight_loss/backend/config"
    "github.com/Afomiat/AI_weight_loss/backend/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers() ([]domain.User, error) {
    collection := config.DB.Collection("users")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var users []domain.User
    for cursor.Next(context.Background()) {
        var user domain.User
        if err = cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func AddUser(user domain.User) error {
    collection := config.DB.Collection("users")
    user.ID = primitive.NewObjectID().Hex()
    _, err := collection.InsertOne(context.Background(), user)
    return err
}
