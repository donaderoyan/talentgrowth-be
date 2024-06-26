package loginhandler

import (
	"net/http"

	login "github.com/donaderoyan/talentgrowth-be/controllers/auth/login"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service login.Service
}

func NewHandlerLogin(service login.Service) *handler {
	return &handler{service: service}
}

// Swagger documentation for LoginHandler
// @Summary User login
// @Description Logs in a user by email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body login.LoginInput true "Login Credentials"
// @Success 200 {object} map[string]interface{} "Login successful, returns access token"
// @Failure 400 {object} map[string]interface{} "Bad request, invalid input"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /login [post]
func (h *handler) LoginHandler(ctx *gin.Context) {
	var input login.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Login failed", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	if errValidator := util.Validator(input, "validate"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPost, errValidator.Error())
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)
	if errLogin != nil {
		switch errLogin.(type) {
		case *login.UserLoginError:
			util.ErrorResponse(ctx, "Login failed", http.StatusInternalServerError, http.MethodPost, errLogin.Error())
			return
		case *login.UserLoginNotFoundError:
			util.ErrorResponse(ctx, "Login failed", http.StatusNotFound, http.MethodPost, errLogin.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}
	} else {
		accessTokenData := map[string]interface{}{
			"id":    resultLogin.ID,
			"email": resultLogin.Email,
		}
		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 24*60) // Expires in 24 hours
		if errToken != nil {
			util.ErrorResponse(ctx, "Failed to generate access token", http.StatusInternalServerError, http.MethodPost, errToken.Error())
			return
		}

		// Remember the token in the context for potential further use in the session
		ctx.Set("accessToken", accessToken)
		// Save the token to the database for session management or token tracking
		errSaveToken := h.service.UpdateRememberTokenService(resultLogin.ID, accessToken)
		if errSaveToken != nil {
			util.ErrorResponse(ctx, "Failed to save access token", http.StatusInternalServerError, http.MethodPost, errSaveToken.Error())
			return
		}

		responseData := map[string]interface{}{
			"id":        resultLogin.ID,
			"email":     resultLogin.Email,
			"firstName": resultLogin.FirstName,
			"lastName":  resultLogin.LastName,
			// "role": resultLogin.Role,
		}

		util.APIResponse(ctx, "Login successful", http.StatusOK, http.MethodPost, map[string]interface{}{
			"accessToken": accessToken,
			"user":        responseData,
		})
	}

}
