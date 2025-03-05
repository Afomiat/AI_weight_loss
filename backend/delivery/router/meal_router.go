package router

import (
	"github.com/Afomiat/AI_weight_loss/backend/config"
	"github.com/Afomiat/AI_weight_loss/backend/delivery/controller"
	"github.com/Afomiat/AI_weight_loss/backend/usecase"
	"github.com/gin-gonic/gin"
)

func MealRoute(env *config.Env, router *gin.RouterGroup) {
    exerciseUsecase := usecase.NewExerciseUsecase()
    exerciseController := controller.NewExerciseController(exerciseUsecase)

    router.GET("/calorie_info", controller.GetCalorieInfo(env))
    router.GET("/meal_suggestion", controller.GetMealSuggestion(env))
    router.POST("/exercise-recommendations", exerciseController.GetExerciseRecommendations)
}

