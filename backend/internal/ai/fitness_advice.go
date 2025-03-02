package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Afomiat/AI_weight_loss/backend/domain"
)


func GetExerciseRecommendation(userGoal string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	url := "https://generativelanguage.googleapis.com/v1/models/gemini-1.5-pro:generateContent?key=" + apiKey

	requestData := domain.GeminiRequest{
		Contents: []domain.Content{
			{
				Role: "user",
				Parts: []domain.Part{
					{Text: "You are a fitness coach recommending exercises for weight loss."},
					{Text: userGoal},
				},
			},
		},
	}

	requestBody, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var geminiResponse domain.GeminiResponse
	json.NewDecoder(resp.Body).Decode(&geminiResponse)

	if len(geminiResponse.Candidates) > 0 {
		return geminiResponse.Candidates[0].Content.Parts[0].Text, nil
	}
	return "No recommendation available", nil
}
