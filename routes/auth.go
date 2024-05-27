package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	login "github.com/donaderoyan/talentgrowth-be/controllers/auth/login"
	register "github.com/donaderoyan/talentgrowth-be/controllers/auth/register"
	loginhandler "github.com/donaderoyan/talentgrowth-be/handlers/auth/login"
	registerhandler "github.com/donaderoyan/talentgrowth-be/handlers/auth/register"
)

func InitAuthRoutes(db *mongo.Database, route *gin.Engine) {

	registerRepository := register.NewRegisterRepository(db)
	registerService := register.NewRegisterService(registerRepository)
	registerHandler := registerhandler.NewHandlerRegister(registerService)

	loginRepository := login.NewLoginRepository(db)
	loginService := login.NewLoginService(loginRepository)
	loginHandler := loginhandler.NewHandlerLogin(loginService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
