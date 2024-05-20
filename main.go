package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	config "github.com/donaderoyan/talentgrowth-be/configs"
	route "github.com/donaderoyan/talentgrowth-be/routes"
	util "github.com/donaderoyan/talentgrowth-be/utils"

	_ "github.com/donaderoyan/talentgrowth-be/docs" // Swagger documentation
)

func main() {
	r := SetupRouter()

	log.Fatal(r.Run(":" + util.GodotEnv("GO_PORT")))
}

// @title Talentgrowth API
// @version 1.0
// @description This is the API documentation for Talentgrowth
// @termsOfService http://swagger.io/terms/
// @contact.name API Support - Donaderoyan
// @contact.email donaderoyan@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func SetupRouter() *gin.Engine {
	mongo := config.ConnectMongoDB()
	db := mongo.Database(util.GodotEnv("MONGO_DBNAME"))
	router := gin.Default()

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	route.InitAuthRoutes(db, router)
	route.InitUserRoutes(db, router)

	return router
}
