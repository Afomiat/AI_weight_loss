package ai

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Afomiat/AI_weight_loss/backend/config"
)

type Meal struct {
    IDMeal          string `json:"idMeal"`
    StrMeal         string `json:"strMeal"`
    StrCategory     string `json:"strCategory"`
    StrArea         string `json:"strArea"`
    StrInstructions string `json:"strInstructions"`
    StrMealThumb    string `json:"strMealThumb"`
    Calories        float64 `json:"calories"`
}

type MealsResponse struct {
    Meals []Meal `json:"meals"`
}

func GetMealSuggestion(dailyCalorieLimit float64) ([]Meal, error) {
    var meals []Meal
    totalCalories := 0.0
    mealTimes := []string{"Breakfast", "Lunch", "Dinner"}

    for range mealTimes {
        meal, err := fetchHealthyMeal(dailyCalorieLimit / 3)
        if err != nil {
            return nil, err
        }
        meals = append(meals, meal)
        totalCalories += meal.Calories
    }

    // Ensure the total calories do not exceed the daily limit
    if totalCalories > dailyCalorieLimit {
        return nil, errors.New("Total calorie intake exceeds daily limit")
    }

    return meals, nil
}

func fetchHealthyMeal(calorieLimit float64) (Meal, error) {
    apiKey := config.GetEnv("MEAL_SUGGESTION_API_KEY")
    url := "https://www.themealdb.com/api/json/v1/" + apiKey + "/random.php"

    req, _ := http.NewRequest("GET", url, nil)
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatalf("Request failed: %v\n", err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var mealsResponse MealsResponse
    err = json.Unmarshal(body, &mealsResponse)
    if err != nil || len(mealsResponse.Meals) == 0 {
        return Meal{}, err
    }

    for _, meal := range mealsResponse.Meals {
        calorieInfo, err := GetCalorieInfo(meal.StrMeal)
        if err != nil || len(calorieInfo.Items) == 0 {
            continue
        }
        meal.Calories = calorieInfo.Items[0].Calories

        if meal.Calories < calorieLimit && meal.Calories > 0 {
            return meal, nil
        }
    }
    return mealsResponse.Meals[0], nil
}
