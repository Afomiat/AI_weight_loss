package usecase

import (
	"github.com/Afomiat/AI_weight_loss/config"
	"github.com/Afomiat/AI_weight_loss/domain"
	"github.com/Afomiat/AI_weight_loss/internal/ai"
	"github.com/Afomiat/AI_weight_loss/repository"
)

func GetMeals() ([]domain.Meal, error) {
    return repository.GetMeals()
}

func GetMealSuggestion(env *config.Env, calorieLimit int) ([]domain.Meal, error) {
    return ai.GetMealSuggestion(env, calorieLimit)
}

func GetCalorieInfo(env *config.Env, food string) (ai.CalorieInfo, error) {
    return ai.GetCalorieInfo(env, food)
}
