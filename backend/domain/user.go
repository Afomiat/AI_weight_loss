package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
    Name      string             `bson:"name" json:"name"`
    Age    int                `bson:"age" json:"age"`
    Goal   string             `bson:"goal" json:"goal"`
    Weight    float64            `bson:"weight" json:"weight"`
	Height    float64            `bson:"height,omitempty" json:"height,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
