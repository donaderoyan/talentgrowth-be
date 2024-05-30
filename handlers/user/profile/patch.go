package profilehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

type handler struct {
	service profile.Service
}

func NewHandlerProfile(service profile.Service) *handler {
	return &handler{service: service}
}

// Swagger documentation for UpdateProfileHandler
// @Summary Update user profile (partial update)
// @Description Update the profile of a user by their ID. Only the fields that are provided will be updated.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body profile.UpdateProfileInput true "Update Profile Data"
// @Success 200 {object} map[string]interface{} "Profile updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/profile/{id} [patch]
func (h *handler) UpdateProfileHandler(ctx *gin.Context) {
	userID := ctx.Param("id") // Assuming 'id' is passed as a URL parameter
	var input profile.UpdateProfileInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}

	if errValidator := util.Validator(input, "updateValidation"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPatch, errValidator.Error())
		return
	}

	updatedUser, errUpdate := h.service.UpdateProfileService(userID, &input)
	if errUpdate != nil {
		switch errUpdate.(type) {
		case *profile.UserProfileUpdateError:
			util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPatch, errUpdate.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPatch, nil)
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

	util.APIResponse(ctx, "Profile updated successfully", http.StatusOK, http.MethodPatch, responseData)
}
