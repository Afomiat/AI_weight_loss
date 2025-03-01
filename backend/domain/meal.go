package domain

type Meal struct {
    ID       string `json:"id" bson:"_id"`
    Name     string `json:"name" bson:"name"`
    Calories int    `json:"calories" bson:"calories"`
}
