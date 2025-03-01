package ai

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Afomiat/AI_weight_loss/backend/config"
)

type CalorieInfo struct {
    Items []struct {
        Name        string  `json:"name"`
        Calories    float64 `json:"calories"`
        Protein     float64 `json:"protein_g"`
        Fat         float64 `json:"fat_total_g"`
        Sugar       float64 `json:"sugar_g"`
    } `json:"items"`
}

func GetCalorieInfo(food string) (CalorieInfo, error) {
    apiKey := config.GetEnv("CALORIE_NINJAS_API_KEY")
    url := "https://api.calorieninjas.com/v1/nutrition?query=" + food

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("X-Api-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatalf("Request failed: %v\n", err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var calorieInfo CalorieInfo
    err = json.Unmarshal(body, &calorieInfo)
    return calorieInfo, err
}
