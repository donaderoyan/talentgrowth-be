package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	config "github.com/donaderoyan/talentgrowth-be/configs"
	route "github.com/donaderoyan/talentgrowth-be/routes"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

func main() {
	r := SetupRouter()

	log.Fatal(r.Run(":" + util.GodotEnv("GO_PORT")))
}

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
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	route.InitAuthRoutes(db, router)
	route.InitUserRoutes(db, router)

	return router
}
