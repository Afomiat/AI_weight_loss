package domain

type Meal struct {
    Name        string   `json:"name"`
    Calories    int      `json:"calories"`
    Ingredients []string `json:"ingredients"`
}
