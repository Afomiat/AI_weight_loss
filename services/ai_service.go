package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Struct to parse API response
type NutritionItem struct {
	Name            string  `json:"name"`
	Calories        float64 `json:"calories"`
	ServingSizeG    float64 `json:"serving_size_g"`
	FatTotalG       float64 `json:"fat_total_g"`
	FatSaturatedG   float64 `json:"fat_saturated_g"`
	ProteinG        float64 `json:"protein_g"`
	CarbohydratesG  float64 `json:"carbohydrates_total_g"`
	FiberG          float64 `json:"fiber_g"`
	SugarG          float64 `json:"sugar_g"`
}

type AIResponse struct {
	Items []NutritionItem `json:"items"`
}

// Function to estimate calories
func EstimateCalories(meal string) (float64, error) {
	apiKey := os.Getenv("CALORIE_NINJAS_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API Key is missing! Set CALORIE_NINJAS_API_KEY")
	}

	// Encode meal properly
	encodedMeal := url.QueryEscape(meal)
	apiURL := fmt.Sprintf("https://api.calorieninjas.com/v1/nutrition?query=%s", encodedMeal)

	// Create request
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("User-Agent", "Mozilla/5.0") // 🔹 Add User-Agent header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read response
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("🔹 Raw API Response:", string(body)) // Log API response

	// Parse JSON
	var result AIResponse
	json.Unmarshal(body, &result)

	// Check if API returned data
	if len(result.Items) > 0 {
		return result.Items[0].Calories, nil
	}

	return 0, fmt.Errorf("no data found for meal: %s", meal)
}
