package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	registerController "github.com/donaderoyan/talentgrowth-be/controllers/auth/register"
	handleRegister "github.com/donaderoyan/talentgrowth-be/handlers/auth/register"
)

func InitAuthRoutes(db *mongo.Database, route *gin.Engine) {

	registerRepository := registerController.NewRegisterRepository(db)
	registerService := registerController.NewRegisterService(registerRepository)
	registerHandler := handleRegister.NewHandlerRegister(registerService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
