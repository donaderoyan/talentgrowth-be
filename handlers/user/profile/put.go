package profilehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

// Swagger documentation for PutProfileHandler
// @Summary Update user profile
// @Description Update the profile of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body profile.UpdateProfileInput true "Update Profile Data"
// @Success 200 {object} map[string]interface{} "Profile updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/profile/{id} [put]
func (h *handler) PutProfileHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, "ID is required")
		return
	}
	var input profile.UpdateProfileInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, err.Error())
		return
	}

	if errValidator := util.Validator(input, "validate"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPut, errValidator.Error())
		return
	}

	updatedUser, errUpdate := h.service.PutProfileService(userID, &input)
	if errUpdate != nil {
		switch errUpdate.(type) {
		case *profile.UserProfileUpdateError:
			util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, errUpdate.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPut, nil)
			return
		}
	}

	responseData := gin.H{
		"firstName":      updatedUser.FirstName,
		"lastName":       updatedUser.LastName,
		"phone":          updatedUser.Phone,
		"address":        updatedUser.Address,
		"birthday":       updatedUser.Birthday,
		"gender":         updatedUser.Gender,
		"nationality":    updatedUser.Nationality,
		"bio":            updatedUser.Bio,
		"profilePicture": updatedUser.ProfilePicture,
	}

	util.APIResponse(ctx, "Profile updated successfully", http.StatusOK, http.MethodPut, responseData)
}
