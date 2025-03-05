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

func GetCalorieInfo(env *config.Env, food string) (CalorieInfo, error) {
    apiKey := env.CalorieNinjasAPIKey 
    url := "https://api.calorieninjas.com/v1/nutrition?query=" + food

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("X-Api-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("Request failed: %v\n", err)
        return CalorieInfo{}, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return CalorieInfo{}, err
    }

    var calorieInfo CalorieInfo
    err = json.Unmarshal(body, &calorieInfo)
    return calorieInfo, err
}
