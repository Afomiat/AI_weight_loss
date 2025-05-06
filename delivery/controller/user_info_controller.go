package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Afomiat/AI_weight_loss/domain"
	"github.com/Afomiat/AI_weight_loss/usecase"
)

type UserInfoController struct {
	UserInfoUsecase usecase.UserInfoUsecase
}

func (c *UserInfoController) AddUserInfo(ctx *gin.Context) {
	var userInfo domain.UserInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := c.UserInfoUsecase.AddUserInfo(ctx, &userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user info"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User info saved successfully"})
}


func (c *UserInfoController) GetUserProgress(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userInfos, err := c.UserInfoUsecase.GetUserProgress(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user progress"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userInfos})
}
