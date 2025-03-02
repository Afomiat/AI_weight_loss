package domain

type WeightLossAdviceRequest struct {
    Weight float64 `json:"weight"`  
    Height float64 `json:"height"`  
    Age    int     `json:"age"`     
    Gender string  `json:"gender"` 
}


type Exercise struct {
    Name     string  `json:"name"`
    Calories float64 `json:"calories"` 
    Duration float64 `json:"duration"` 
}

type WeightLossAdviceResponse struct {
    BMI         float64   `json:"bmi"`
    IdealWeight float64   `json:"ideal_weight"` 
    Exercises   []Exercise `json:"exercises"`
    Message     string    `json:"message"`
}