package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	musicalinfo "github.com/donaderoyan/talentgrowth-be/controllers/user/musicalinfo"
	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	musicalinfohandler "github.com/donaderoyan/talentgrowth-be/handlers/user/musicalinfo"
	profilehandler "github.com/donaderoyan/talentgrowth-be/handlers/user/profile"
	middleware "github.com/donaderoyan/talentgrowth-be/middlewares"
)

func InitUserRoutes(db *mongo.Database, route *gin.Engine) {
	// Initialize profile
	profileRepository := profile.NewProfileRepository(db)
	profileService := profile.NewProfileService(profileRepository)
	profileHandler := profilehandler.NewHandlerProfile(profileService)

	// Initialize musical information
	musicalInfoRepository := musicalinfo.NewMusicalInfoRepository(db)
	musicalInfoService := musicalinfo.NewMusicalInfoService(musicalInfoRepository)
	musicalInfoHandler := musicalinfohandler.NewMusicalInfohandler(musicalInfoService)

	userGroup := route.Group("/api/v1/user").Use(middleware.Auth(db))
	// profile
	userGroup.PATCH("/profile/:id", profileHandler.UpdateProfileHandler)
	userGroup.PUT("/profile/:id", profileHandler.PutProfileHandler)
	userGroup.GET("/profile/:id", profileHandler.GetProfileHandler)
	// musical information
	userGroup.PATCH("/musicalinfo/:id", musicalInfoHandler.UpdateMusicalInfoHandler)
}
