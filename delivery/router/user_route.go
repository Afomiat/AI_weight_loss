package router

import (
	"time"

	"github.com/Afomiat/AI_weight_loss/config"
	"github.com/Afomiat/AI_weight_loss/delivery/controller"
	"github.com/Afomiat/AI_weight_loss/repository"
	"github.com/Afomiat/AI_weight_loss/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserInfoRouter(env *config.Env, timeout time.Duration, db *mongo.Database, router *gin.RouterGroup) {
	userInfoRepo := repository.NewUserInfoRepository(db, "user_info")

	userInfoUsecase := usecase.NewUserInfoUsecase(userInfoRepo, timeout)
	userInfoController := &controller.UserInfoController{
		UserInfoUsecase: userInfoUsecase,
	}

	router.POST("/user_info", userInfoController.AddUserInfo)
	router.GET("/user_info/:user_id", userInfoController.GetUserProgress)
}
