package util

import "github.com/gin-gonic/gin"

type ResponsesEntity struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponseEntity struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

func APIResponse(ctx *gin.Context, message string, statusCode int, method string, data interface{}) {
	jsonResponse := ResponsesEntity{
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

func ErrorResponse(ctx *gin.Context, message string, statusCode int, method string, err interface{}) {
	errorResponse := ErrorResponseEntity{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Error:      err,
	}

	ctx.JSON(statusCode, errorResponse)
	ctx.AbortWithStatus(statusCode)
}
