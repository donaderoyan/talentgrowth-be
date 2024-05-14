package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	loginController "github.com/donaderoyan/talentgrowth-be/controllers/auth/login"
	registerController "github.com/donaderoyan/talentgrowth-be/controllers/auth/register"
	handlerLogin "github.com/donaderoyan/talentgrowth-be/handlers/auth/login"
	handlerRegister "github.com/donaderoyan/talentgrowth-be/handlers/auth/register"
)

func InitAuthRoutes(db *mongo.Database, route *gin.Engine) {

	registerRepository := registerController.NewRegisterRepository(db)
	registerService := registerController.NewRegisterService(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	loginRepository := loginController.NewLoginRepository(db)
	loginService := loginController.NewLoginService(loginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
