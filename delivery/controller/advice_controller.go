package controller

import (
	"net/http"

	"github.com/Afomiat/AI_weight_loss/usecase"
	"github.com/gin-gonic/gin"
)

type ExerciseController struct {
	Usecase *usecase.ExerciseUsecase
}

func NewExerciseController(usecase *usecase.ExerciseUsecase) *ExerciseController {
	return &ExerciseController{Usecase: usecase}
}

func (ec *ExerciseController) GetExerciseRecommendations(c *gin.Context) {
	var request struct {
		Goal string `json:"goal"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	recommendation, err := ec.Usecase.GetRecommendation(request.Goal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recommendation": recommendation})
}
