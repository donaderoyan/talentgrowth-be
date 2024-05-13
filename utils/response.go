package util

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

func APIResponse(ctx *gin.Context, message string, statusCode int, method string, data interface{}) {
	jsonResponse := Responses{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Data:       data,
	}

	ctx.JSON(statusCode, jsonResponse)
	if statusCode >= 400 {
		ctx.AbortWithStatus(statusCode)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, message string, statusCode int, method string, err interface{}) {
	errorResponse := ErrorResponse{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Error:      err,
	}

	ctx.JSON(statusCode, errorResponse)
	ctx.AbortWithStatus(statusCode)
}
