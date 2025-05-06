package router

import (
	"time"

	"github.com/Afomiat/AI_weight_loss/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(r *gin.Engine, env *config.Env, db *mongo.Database) *gin.Engine {

	api := r.Group("/api") 

	MealRoute(env, api)           
	NewUserInfoRouter(env, 10*time.Second, db, api) 

	return r
}
