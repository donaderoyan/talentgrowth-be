package registerhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	register "github.com/donaderoyan/talentgrowth-be/controllers/auth/register"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

// Swagger documentation for Register API
// @Summary Register new user
// @Description Register a new user with email, password, first name, and last name
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body register.RegisterInput true "Register Input"
// @Success 200 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad request, invalid input"
// @Failure 409 {object} map[string]interface{} "Conflict, user already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/register [post]
func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Register new account failed", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	if errValidator := util.Validator(input, "validate"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPost, errValidator.Error())
		return
	}

	responseData, errRegister := h.service.RegisterService(&input)
	if errRegister != nil {
		switch errRegister.(type) {
		case *register.UserAlreadyExistsError:
			util.ErrorResponse(ctx, "Register new account failed", http.StatusBadRequest, http.MethodPost, errRegister.Error())
			return
		case *register.UserRegistrationError:
			util.ErrorResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, errRegister.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

	}

	input.FirstName = responseData.FirstName
	input.LastName = responseData.LastName
	input.Email = responseData.Email
	input.Password = "****" // Mask the password for security reasons

	util.APIResponse(ctx, "Register new account success", http.StatusOK, http.MethodPost, input)
}
