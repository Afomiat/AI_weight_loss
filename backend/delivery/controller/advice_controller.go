package controller

import (
	"net/http"

	"github.com/Afomiat/AI_weight_loss/backend/usecase"
	"github.com/gin-gonic/gin"
)

func GetAdvice(c *gin.Context) {
    advice, err := usecase.GetAdvice()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, advice)
}
