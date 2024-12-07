package httphelper

import (
	"net/http"

	"github.com/Shachindra/altlife/server/api/types"

	"github.com/gin-gonic/gin"
)

// TODO: add method for internet error with common msg
func ErrResponse(c *gin.Context, statusCode int, errMessage string) {
	response := types.ApiResponse{
		Status: statusCode,
		Error:  errMessage,
	}
	c.JSON(response.Status, response)
}

func CErrResponse(c *gin.Context, status int, customStatusCode int, errMessage string) {
	response := types.ApiResponse{
		Status: customStatusCode,
		Error:  errMessage,
	}
	c.JSON(status, response)
}

func SuccessResponse(c *gin.Context, message string, payload interface{}) {
	response := types.ApiResponse{
		Status:  http.StatusOK,
		Payload: payload,
		Result:  message,
	}
	c.JSON(response.Status, response)
}

func InternalServerError(c *gin.Context) {
	response := types.ApiResponse{
		Status: http.StatusInternalServerError,
		Error:  "unexpected error occurred",
	}
	c.JSON(response.Status, response)
}

func NewInternalServerError(c *gin.Context, message string, format string, args ...interface{}) {
	response := types.ApiResponse{
		Status: http.StatusInternalServerError,
		Error:  "unexpected error occurred",
		Result: message,
	}
	c.JSON(response.Status, response)
}
