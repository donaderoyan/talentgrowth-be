package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	profileController "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	handlerProfile "github.com/donaderoyan/talentgrowth-be/handlers/user/profile"
)

func InitUserRoutes(db *mongo.Database, route *gin.Engine) {
	profileRepository := profileController.NewProfileRepository(db)
	profileService := profileController.NewProfileService(profileRepository)
	profileHandler := handlerProfile.NewHandlerProfile(profileService)

	userGroup := route.Group("/api/v1/user")
	userGroup.PUT("/profile/:id", profileHandler.UpdateProfileHandler)
}
