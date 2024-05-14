package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dgrijalva/jwt-go"
	model "github.com/donaderoyan/talentgrowth-be/models"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/gin-gonic/gin"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errorResponse UnathorizatedError
		errorResponse.Status = "Unauthorized"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "accessToken invalid or expired"

		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			errorResponse.Message = "Authorization header is missing"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		splitToken := strings.Split(tokenHeader, "Bearer ")
		if len(splitToken) != 2 {
			errorResponse.Message = "Malformed Authorization token"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		accessToken := strings.TrimSpace(splitToken[1])

		// Verify the token
		token, err := util.VerifyToken(accessToken, "JWT_SECRET")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errorResponse.Message = "Error parsing token claims"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		userId, ok := claims["id"].(string)
		if !ok {
			errorResponse.Message = "User ID not found in token claims"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		var user model.User
		userIDPrimitive, _ := primitive.ObjectIDFromHex(userId)
		if err := db.Collection("users").FindOne(ctx, bson.M{"_id": userIDPrimitive}).Decode(&user); err != nil {
			errorResponse.Message = "User not found"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		if user.RememberToken != accessToken {
			errorResponse.Message = fmt.Sprintf("Token does not match stored token remembertoken: %s | userId: %s", user.RememberToken, user)
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		}

		// Set user information in context
		ctx.Set("user", user)
		ctx.Next()
	}
}
