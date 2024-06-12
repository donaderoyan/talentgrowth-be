package profilehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	profile "github.com/donaderoyan/talentgrowth-be/controllers/user/profile"
	util "github.com/donaderoyan/talentgrowth-be/utils"
)

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
// @Router /user/profile/{id} [patch]
func (h *handler) UpdateProfileHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPut, "ID is required")
		return
	}
	var input bson.M

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Update profile failed", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}

	var addressMap bson.M
	var validateAddress profile.Address
	if address, ok := input["address"]; ok {
		addressMap, _ = address.(map[string]interface{})

		validateAddress = profile.Address{
			Street:     util.GetStringFromMap(addressMap, "street"),
			City:       util.GetStringFromMap(addressMap, "city"),
			State:      util.GetStringFromMap(addressMap, "state"),
			Country:    util.GetStringFromMap(addressMap, "country"),
			PostalCode: util.GetStringFromMap(addressMap, "postalCode"),
		}
	}

	validateInput := profile.UpdateProfileInput{
		FirstName:      util.GetStringFromMap(input, "firstName"),
		LastName:       util.GetStringFromMap(input, "lastName"),
		Phone:          util.GetStringFromMap(input, "phone"),
		Birthday:       util.GetStringFromMap(input, "birthday"),
		Gender:         util.GetStringFromMap(input, "gender"),
		Nationality:    util.GetStringFromMap(input, "nationality"),
		Bio:            util.GetStringFromMap(input, "bio"),
		ProfilePicture: util.GetStringFromMap(input, "profilePicture"),
		Address:        validateAddress,
	}

	if errValidator := util.Validator(validateInput, "updateValidation"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPatch, errValidator.Error())
		return
	}

	updatedUser, errUpdate := h.service.PatchProfileService(userID, input)
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

	util.APIResponse(ctx, "Profile updated successfully", http.StatusOK, http.MethodPatch, responseData)
}
