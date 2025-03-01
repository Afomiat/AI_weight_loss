package ai

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Afomiat/AI_weight_loss/backend/config"
	"github.com/Afomiat/AI_weight_loss/backend/domain"
)

type SpoonacularMealSuggestion struct {
	Meals []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"meals"`
}

type SpoonacularRecipeInformation struct {
	ExtendedIngredients []struct {
		Name string  `json:"name"`
	} `json:"extendedIngredients"`
}

type CalorieNinjaResponse struct {
	Items []struct {
		Name     string  `json:"name"`
		Calories float64 `json:"calories"`
	} `json:"items"`
}

func GetMealSuggestion(calorieLimit int) ([]domain.Meal, error) {
	
	spoonacularAPIKey := config.GetEnv("SPOONACULAR_API_KEY")
	if spoonacularAPIKey == "" {
		return nil, errors.New("Spoonacular API key not found")
	}

	spoonacularURL := "https://api.spoonacular.com/mealplanner/generate?timeFrame=day&targetCalories=" + strconv.Itoa(calorieLimit)

	spoonacularReq, _ := http.NewRequest("GET", spoonacularURL, nil)
	spoonacularReq.Header.Add("x-api-key", spoonacularAPIKey)

	spoonacularRes, err := http.DefaultClient.Do(spoonacularReq)
	if err != nil {
		log.Fatalf("Spoonacular API request failed: %v\n", err)
		return nil, err
	}
	defer spoonacularRes.Body.Close()

	if spoonacularRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch meal suggestions from Spoonacular API")
	}

	var spoonacularResponse SpoonacularMealSuggestion
	spoonacularBody, _ := ioutil.ReadAll(spoonacularRes.Body)
	err = json.Unmarshal(spoonacularBody, &spoonacularResponse)
	if err != nil {
		return nil, err
	}

	
	calorieNinjaAPIKey := config.GetEnv("CALORIE_NINJAS_API_KEY")
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
			log.Fatalf("Spoonacular recipe information request failed: %v\n", err)
			return nil, err
		}
		defer recipeRes.Body.Close()

		if recipeRes.StatusCode != http.StatusOK {
			log.Printf("Failed to fetch recipe information for meal: %s\n", meal.Title)
			continue 
		}

		var recipeInformation SpoonacularRecipeInformation
		recipeBody, _ := ioutil.ReadAll(recipeRes.Body)
		err = json.Unmarshal(recipeBody, &recipeInformation)
		if err != nil {
			log.Printf("Failed to parse recipe information for meal: %s\n", meal.Title)
			continue 
		}

		mealCalories := 0
		for _, ingredient := range recipeInformation.ExtendedIngredients {
			calorieNinjaURL := "https://api.calorieninjas.com/v1/nutrition?query=" + ingredient.Name

			calorieNinjaReq, _ := http.NewRequest("GET", calorieNinjaURL, nil)
			calorieNinjaReq.Header.Add("X-Api-Key", calorieNinjaAPIKey)

			calorieNinjaRes, err := http.DefaultClient.Do(calorieNinjaReq)
			if err != nil {
				log.Fatalf("Calorie Ninja API request failed: %v\n", err)
				return nil, err
			}
			defer calorieNinjaRes.Body.Close()

			if calorieNinjaRes.StatusCode != http.StatusOK {
				log.Printf("Failed to fetch calorie information for ingredient: %s\n", ingredient.Name)
				continue
			}

			var calorieNinjaResponse CalorieNinjaResponse
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
		}

		if totalCalories+mealCalories > calorieLimit {
			break 
		}

		meals = append(meals, domain.Meal{
			Name:     meal.Title,
			Calories: mealCalories,
		})
		totalCalories += mealCalories

		if len(meals) >= 3 { 
			break
		}
	}

	if len(meals) == 0 {
		return nil, errors.New("no meals found within the calorie limit")
	}

	return meals, nil
}