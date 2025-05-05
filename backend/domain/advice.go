package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeightLossAdviceRequest struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
}

type Exercise struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	Name     string  `json:"name"`
	Calories float64 `json:"calories"`
	Duration float64 `json:"duration"`
	Goal     string  `bson:"goal" json:"goal"`
	Advice   string  `bson:"advice" json:"advice"`
}

type WeightLossAdviceResponse struct {
	BMI         float64    `json:"bmi"`
	IdealWeight float64    `json:"ideal_weight"`
	Exercises   []Exercise `json:"exercises"`
	Message     string     `json:"message"`
}