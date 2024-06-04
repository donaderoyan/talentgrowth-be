package musicalinfohandler

import (
	"net/http"

	musicalinfo "github.com/donaderoyan/talentgrowth-be/controllers/user/musicalinfo"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateMusicalInfoHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	var input musicalinfo.MusicalInfoInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.ErrorResponse(ctx, "Create musical information failed", http.StatusBadRequest, http.MethodPost, err.Error())
	}

	if errValidator := util.Validator(input, "validate"); errValidator != nil {
		util.ErrorResponse(ctx, "The input value is invalid", http.StatusBadRequest, http.MethodPost, errValidator.Error())
		return
	}

	updateMusicalInfo, errUpdate := h.service.CreateMusicalInfoService(userID, &input)
	if errUpdate != nil {
		switch errUpdate.(type) {
		case *musicalinfo.MusicalInfoCreateError:
			util.ErrorResponse(ctx, "Create musical information failed", http.StatusBadRequest, http.MethodPost, errUpdate.Error())
			return
		default:
			// Handle other unexpected errors
			util.ErrorResponse(ctx, "Internal server error", http.StatusInternalServerError, http.MethodPost, errUpdate.Error())
			return
		}
	}

	util.APIResponse(ctx, "Musical information created successfully", http.StatusOK, http.MethodPost, updateMusicalInfo)
}
