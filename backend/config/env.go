package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env struct to store environment variables
type Env struct {
	ServerPort             string
	MongoURI               string
	DBName                 string
	CalorieNinjasAPIKey    string
	MealSuggestionAPIKey   string
	SpoonacularAPIKey      string
	GeminiAPIKey           string
}

// LoadEnv loads environment variables from .env file
func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := &Env{
		ServerPort:           os.Getenv("SERVER_PORT"),
		MongoURI:             os.Getenv("MONGO_URI"),
		DBName:               os.Getenv("DB_NAME"),
		CalorieNinjasAPIKey:  os.Getenv("CALORIE_NINJAS_API_KEY"),
		MealSuggestionAPIKey: os.Getenv("MEAL_SUGGESTION_API_KEY"),
		SpoonacularAPIKey:    os.Getenv("SPOONACULAR_API_KEY"),
		GeminiAPIKey:         os.Getenv("GEMINI_API_KEY"),
	}

	return env
}
