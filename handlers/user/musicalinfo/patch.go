package musicalinfohandler

import (
	"net/http"

	musicalinfo "github.com/donaderoyan/talentgrowth-be/controllers/user/musicalinfo"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Swagger documentation for UpdateMusicalInfoHandler
// @Summary Update musical information (partial update)
// @Description Update musical information for a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body musicalinfo.MusicalInfoInput true "Musical information to update"
// @Success 200 {object} map[string]interface{} "Musical information updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/musicalinfo/{id} [patch]
func (h *handler) UpdateMusicalInfoHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	var input bson.M

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Update musical information failed", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}

	secondaryInstrumentsSlice, ok := input["secondaryInstruments"].([]interface{})
	if !ok {
		secondaryInstrumentsSlice = nil
	}
	genresSlice, ok := input["genres"].([]interface{})
	if !ok {
		genresSlice = nil
	}
	favoriteArtistsSlice, ok := input["favoriteArtists"].([]interface{})
	if !ok {
		favoriteArtistsSlice = nil
	}
	learningGoalsSlice, ok := input["learningGoals"].([]interface{})
	if !ok {
		learningGoalsSlice = nil
	}

	validateInput := musicalinfo.MusicalInfoInput{
		SkillLevel:           input["skillLevel"].(string),
		PrimaryInstrument:    input["primaryInstrument"].(string),
		SecondaryInstruments: util.InterfaceSliceToStringSlice(secondaryInstrumentsSlice),
		Genres:               util.InterfaceSliceToStringSlice(genresSlice),
		FavoriteArtists:      util.InterfaceSliceToStringSlice(favoriteArtistsSlice),
		LearningGoals:        util.InterfaceSliceToStringSlice(learningGoalsSlice),
	}
	if errValidator := util.Validator(validateInput, "updateValidation"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPatch, errValidator.Error())
		return
	}

	updateMusicalInfo, errUpdate := h.service.UpdateMusicalInfoService(userID, input)
	if errUpdate != nil {
		switch errUpdate.(type) {
		case *musicalinfo.MusicalInfoUpdateError:
			util.ErrorResponse(ctx, "Update musical information failed", http.StatusBadRequest, http.MethodPatch, errUpdate.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPatch, nil)
			return
		}
	}

	util.APIResponse(ctx, "Musical information updated successfully", http.StatusOK, http.MethodPatch, updateMusicalInfo)
}
