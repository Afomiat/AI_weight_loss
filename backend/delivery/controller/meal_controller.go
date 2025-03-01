package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Afomiat/AI_weight_loss/backend/usecase"
)

func GetMeals(c *gin.Context) {
    meals, err := usecase.GetMeals()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, meals)
}

// func GetMealSuggestion(c *gin.Context) {
//     dailyCalorieLimitStr := c.Query("dailyCalorieLimit")
//     dailyCalorieLimit, err := strconv.ParseFloat(dailyCalorieLimitStr, 64)
//     if err != nil || dailyCalorieLimit <= 0 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid daily calorie limit"})
//         return
//     }

//     suggestions, err := usecase.GetMealSuggestion(dailyCalorieLimit)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     c.JSON(http.StatusOK, suggestions)
// }

func GetCalorieInfo(c *gin.Context) {
    food := c.Query("food")
    if food == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Food query parameter is required"})
        return
    }
    calorieInfo, err := usecase.GetCalorieInfo(food)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, calorieInfo)
}
