package ai

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Afomiat/AI_weight_loss/config"
	"github.com/Afomiat/AI_weight_loss/domain"
)

func GetMealSuggestion(env *config.Env, calorieLimit int) ([]domain.Meal, error) {
	spoonacularAPIKey := env.SpoonacularAPIKey
	if spoonacularAPIKey == "" {
		return nil, errors.New("Spoonacular API key not found")
	}

	spoonacularURL := "https://api.spoonacular.com/mealplanner/generate?timeFrame=day&targetCalories=" + strconv.Itoa(calorieLimit)
	spoonacularReq, _ := http.NewRequest("GET", spoonacularURL, nil)
	spoonacularReq.Header.Add("x-api-key", spoonacularAPIKey)

	spoonacularRes, err := http.DefaultClient.Do(spoonacularReq)
	if err != nil {
		log.Printf("Spoonacular API request failed: %v\n", err)
		return nil, err
	}
	defer spoonacularRes.Body.Close()

	if spoonacularRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch meal suggestions from Spoonacular API")
	}

	var spoonacularResponse domain.SpoonacularMealSuggestion
	spoonacularBody, _ := ioutil.ReadAll(spoonacularRes.Body)
	err = json.Unmarshal(spoonacularBody, &spoonacularResponse)
	if err != nil {
		return nil, err
	}

	calorieNinjaAPIKey := env.CalorieNinjasAPIKey
	if calorieNinjaAPIKey == "" {
		return nil, errors.New("Calorie Ninja API key not found")
	}

	var meals []domain.Meal
	totalCalories := 0

	for _, meal := range spoonacularResponse.Meals {
		recipeURL := "https://api.spoonacular.com/recipes/" + strconv.Itoa(meal.ID) + "/information?includeNutrition=false"
		recipeReq, _ := http.NewRequest("GET", recipeURL, nil)
		recipeReq.Header.Add("x-api-key", spoonacularAPIKey)

		recipeRes, err := http.DefaultClient.Do(recipeReq)
		if err != nil {
			log.Printf("Spoonacular recipe request failed: %v\n", err)
			continue
		}
		defer recipeRes.Body.Close()

		if recipeRes.StatusCode != http.StatusOK {
			log.Printf("Failed to fetch recipe info for meal: %s\n", meal.Title)
			continue
		}

		var recipeInformation domain.SpoonacularRecipeInformation
		recipeBody, _ := ioutil.ReadAll(recipeRes.Body)
		err = json.Unmarshal(recipeBody, &recipeInformation)
		if err != nil {
			log.Printf("Failed to parse recipe info for meal: %s\n", meal.Title)
			continue
		}

		mealCalories := 0
		var ingredientsList []string

		for _, ingredient := range recipeInformation.ExtendedIngredients {
			calorieNinjaURL := "https://api.calorieninjas.com/v1/nutrition?query=" + ingredient.Name
			calorieNinjaReq, _ := http.NewRequest("GET", calorieNinjaURL, nil)
			calorieNinjaReq.Header.Add("X-Api-Key", calorieNinjaAPIKey)

			calorieNinjaRes, err := http.DefaultClient.Do(calorieNinjaReq)
			if err != nil {
				log.Printf("Calorie Ninja API request failed: %v\n", err)
				continue
			}
			defer calorieNinjaRes.Body.Close()

			if calorieNinjaRes.StatusCode != http.StatusOK {
				log.Printf("Failed to fetch calorie info for ingredient: %s\n", ingredient.Name)
				continue
			}

			var calorieNinjaResponse domain.CalorieNinjaResponse
			calorieNinjaBody, _ := ioutil.ReadAll(calorieNinjaRes.Body)
			err = json.Unmarshal(calorieNinjaBody, &calorieNinjaResponse)
			if err != nil {
				log.Printf("Failed to parse Calorie Ninja API response for ingredient: %s\n", ingredient.Name)
				continue
			}

			if len(calorieNinjaResponse.Items) == 0 {
				log.Printf("No calorie data found for ingredient: %s\n", ingredient.Name)
				continue
			}

			mealCalories += int(calorieNinjaResponse.Items[0].Calories)
			ingredientsList = append(ingredientsList, ingredient.Name)
		}

		if totalCalories+mealCalories > calorieLimit {
			break
		}

		meals = append(meals, domain.Meal{
			Name:        meal.Title,
			Calories:    mealCalories,
			Ingredients: ingredientsList,
		})
		totalCalories += mealCalories
	}

	if len(meals) == 0 {
		return nil, errors.New("no meals found within the calorie limit")
	}

	return meals, nil
}
