package profilehandler

import (
	"net/http"

	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/gin-gonic/gin"
)

// Swagger documentation for GetProfileHandler
// @Summary Get user profile
// @Description Get the profile of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} profile.UpdateProfileInput
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/profile/{id} [get]
func (h *handler) GetProfileHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodGet, "ID is required")
		return
	}

	dataUser, errData := h.service.GetProfileService(userID)
	if errData != nil {
		switch errData.(type) {
		case *profile.GetUserProfileError:
			util.ErrorResponse(ctx, "Get user profile failed", http.StatusBadRequest, http.MethodGet, errData.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodGet, nil)
			return
		}
	}

	// masking data
	responseData := gin.H{
		"firstName":   dataUser.FirstName,
		"lastName":    dataUser.LastName,
		"phone":       dataUser.Phone,
		"address":     dataUser.Address,
		"birthday":    dataUser.Birthday,
		"gender":      dataUser.Gender,
		"nationality": dataUser.Nationality,
		"bio":         dataUser.Bio,
	}

	util.APIResponse(ctx, "Get user profile successfully", http.StatusOK, http.MethodGet, responseData)
}
