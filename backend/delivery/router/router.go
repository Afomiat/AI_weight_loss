package router

import (
	"github.com/Afomiat/AI_weight_loss/backend/delivery/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/meals", controller.GetMeals)
    r.GET("/meal-suggestion", controller.GetMealSuggestion)
    r.GET("/calories", controller.GetCalorieInfo)
    // r.GET("/advice", controller.GetAdvice)
    return r
}
