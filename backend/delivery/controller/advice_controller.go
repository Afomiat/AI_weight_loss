package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Afomiat/AI_weight_loss/backend/usecase"
)

func GetAdvice(c *gin.Context) {
    advice, err := usecase.GetAdvice()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, advice)
}
