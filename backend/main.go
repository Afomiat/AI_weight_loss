package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Afomiat/AI_weight_loss/services"
	"github.com/gin-gonic/gin"
)

type MealRequest struct {
	Meal string `json:"meal"`
}

type MealResponse struct {
	Meal     string  `json:"meal"`
	Calories float64 `json:"calories"`
}

func EstimateCaloriesHandler(c *gin.Context) {
	var req MealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	calories, err := services.EstimateCalories(req.Meal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MealResponse{
		Meal:     req.Meal,
		Calories: calories,
	})
}

func main() {
	router := gin.Default()

	router.POST("/estimate_calories", EstimateCaloriesHandler)

	
	fmt.Println("🚀 Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
