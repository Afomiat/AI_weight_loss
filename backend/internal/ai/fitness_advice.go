package ai

import (
    "net/http"
    "io/ioutil"
    "log"
    "github.com/Afomiat/AI_weight_loss/backend/config"
)

func GetFitnessAdvice() (string, error) {
    apiKey := config.GetEnv("EXERCISEDB_API_KEY")
    url := "https://exercisedb.p.rapidapi.com/status"

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("X-Api-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatalf("Request failed: %v\n", err)
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    return string(body), nil
}
