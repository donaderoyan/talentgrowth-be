package handleRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"

	registerController "github.com/donaderoyan/talentgrowth-be/controllers/auth/register"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

type handler struct {
	service registerController.Service
}

func NewHandlerRegister(service registerController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input registerController.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ValidatorErrorResponse(ctx, "Register new account failed", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	if errValidator := util.Validator(input); errValidator != nil {
		util.ValidatorErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPost, errValidator.Error())
		return
	}

	_, errRegister := h.service.RegisterService(&input)
	if errRegister != nil {
		switch errRegister.(type) {
		case *registerController.UserAlreadyExistsError:
			util.ValidatorErrorResponse(ctx, "Register new account failed", http.StatusBadRequest, http.MethodPost, errRegister.Error())
			return
		case *registerController.UserRegistrationError:
			util.ValidatorErrorResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, errRegister.Error())
			return
		default:
			// Handle other unexpected errors
			util.ValidatorErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

	}

	util.APIResponse(ctx, "Register new account success", http.StatusOK, http.MethodPost, nil)
}
