package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	profilehandler "github.com/donaderoyan/talentgrowth-be/handlers/user/profile"
	middleware "github.com/donaderoyan/talentgrowth-be/middlewares"
)

func InitUserRoutes(db *mongo.Database, route *gin.Engine) {
	profileRepository := profile.NewProfileRepository(db)
	profileService := profile.NewProfileService(profileRepository)
	profileHandler := profilehandler.NewHandlerProfile(profileService)

	userGroup := route.Group("/api/v1/user").Use(middleware.Auth(db))
	userGroup.PATCH("/profile/:id", profileHandler.UpdateProfileHandler)
}
