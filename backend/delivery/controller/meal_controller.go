package controller

import (
	"net/http"
	"strconv"

	"github.com/Afomiat/AI_weight_loss/backend/usecase"
	"github.com/gin-gonic/gin"
)

func GetMeals(c *gin.Context) {
    meals, err := usecase.GetMeals()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, meals)
}



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

func GetMealSuggestion(c *gin.Context) {
	calorieLimit := c.Query("calorie_limit")
	if calorieLimit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Calorie limit query parameter is required"})
		return
	}

	limit, err := strconv.Atoi(calorieLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid calorie limit"})
		return
	}

	meals, err := usecase.GetMealSuggestion(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meals)
}