package domain

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

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Role    string `json:"role"`
	Parts   []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content Content `json:"content"`
	} `json:"candidates"`
}
