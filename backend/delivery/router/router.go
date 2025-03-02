package router

import (
	"github.com/Afomiat/AI_weight_loss/backend/delivery/controller"
	"github.com/Afomiat/AI_weight_loss/backend/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    exerciseUsecase := usecase.NewExerciseUsecase()
    exerciseController := controller.NewExerciseController(exerciseUsecase)

    r.GET("/meals", controller.GetMeals)
    r.GET("/meal-suggestion", controller.GetMealSuggestion)
    r.GET("/calories", controller.GetCalorieInfo)
    r.POST("/exercise-recommendations", exerciseController.GetExerciseRecommendations)

    return r
}
