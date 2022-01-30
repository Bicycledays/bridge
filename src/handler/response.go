package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type resultResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func newResultResponse(c *gin.Context, result interface{}) {
	log.Println("response", result)
	c.JSON(http.StatusOK, resultResponse{
		Success: true,
		Result:  result,
	})
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, message, error string) {
	log.Println("error", message, error)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Success: false,
		Message: message,
		Error:   error,
	})
}
