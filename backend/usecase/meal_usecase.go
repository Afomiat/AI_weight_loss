package usecase

import (
    "github.com/Afomiat/AI_weight_loss/backend/domain"
    "github.com/Afomiat/AI_weight_loss/backend/repository"
    "github.com/Afomiat/AI_weight_loss/backend/internal/ai"
)

func GetMeals() ([]domain.Meal, error) {
    return repository.GetMeals()
}

// func GetMealSuggestion(dailyCalorieLimit float64) ([]domain.Meal, error) {
//     meals, err := ai.GetMealSuggestion(dailyCalorieLimit)
//     if err != nil {
//         return nil, err
//     }

//     var domainMeals []domain.Meal
//     for _, meal := range meals {
//         domainMeals = append(domainMeals, domain.Meal{
//             ID:       meal.IDMeal,
//             Name:     meal.StrMeal,
//             Calories: int(meal.Calories),
//         })
//     }

//     return domainMeals, nil
// }

func GetCalorieInfo(food string) (ai.CalorieInfo, error) {
    return ai.GetCalorieInfo(food)
}
