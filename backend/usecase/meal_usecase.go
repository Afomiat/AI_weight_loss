package usecase

import (
    "github.com/Afomiat/AI_weight_loss/backend/domain"
    "github.com/Afomiat/AI_weight_loss/backend/repository"
    "github.com/Afomiat/AI_weight_loss/backend/internal/ai"
)

func GetMeals() ([]domain.Meal, error) {
    return repository.GetMeals()
}

func GetCalorieInfo(food string) (ai.CalorieInfo, error) {
    return ai.GetCalorieInfo(food)
}
func GetMealSuggestion(calorieLimit int) ([]domain.Meal, error) {
	return ai.GetMealSuggestion(calorieLimit)
}