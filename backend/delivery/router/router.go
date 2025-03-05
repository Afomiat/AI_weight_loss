package router

import (
	"time"

	"github.com/Afomiat/AI_weight_loss/backend/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(env *config.Env, db *mongo.Database) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api") 

	MealRoute(env, api)           
	NewUserInfoRouter(env, 10*time.Second, db, api) 

	return r
}
