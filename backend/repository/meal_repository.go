package repository

import (
    "context"
    "github.com/Afomiat/AI_weight_loss/backend/config"
    "github.com/Afomiat/AI_weight_loss/backend/domain"
    "go.mongodb.org/mongo-driver/bson"
)

func GetMeals() ([]domain.Meal, error) {
    collection := config.DB.Collection("meals")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var meals []domain.Meal
    for cursor.Next(context.Background()) {
        var meal domain.Meal
        if err = cursor.Decode(&meal); err != nil {
            return nil, err
        }
        meals = append(meals, meal)
    }
    return meals, nil
}

func AddMeal(meal domain.Meal) error {
    collection := config.DB.Collection("meals")
    // meal. = primitive.NewObjectID().Hex()
    _, err := collection.InsertOne(context.Background(), meal)
    return err
}
