package handlerProfile

import (
	"net/http"

	"github.com/gin-gonic/gin"

	profileController "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

type handler struct {
	service profileController.Service
}

func NewHandlerProfile(service profileController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateProfileHandler(ctx *gin.Context) {
	userID := ctx.Param("id") // Assuming 'id' is passed as a URL parameter
	var input profileController.UpdateProfileInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, err.Error())
		return
	}

	if errValidator := util.Validator(input); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPut, errValidator.Error())
		return
	}

	// valid, err := profileController.ValidateBirthday(input.Birthday)
	// if !valid {
	// 	util.ErrorResponse(ctx, "Invalid birthday format", http.StatusBadRequest, http.MethodPut, err.Error())
	// 	return
	// }

	updatedUser, errUpdate := h.service.UpdateProfileService(userID, &input)
	if errUpdate != nil {
		switch errUpdate.(type) {
		case *profileController.UserProfileUpdateError:
			util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, errUpdate.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPut, errUpdate.Error())
			return
		}
	}

	responseData := gin.H{
		"first_name":  updatedUser.FirstName,
		"last_name":   updatedUser.LastName,
		"phone":       updatedUser.Phone,
		"address":     updatedUser.Address,
		"birthday":    updatedUser.Birthday,
		"gender":      updatedUser.Gender,
		"nationality": updatedUser.Nationality,
		"bio":         updatedUser.Bio,
	}

	util.APIResponse(ctx, "Profile updated successfully", http.StatusOK, http.MethodPut, responseData)
}