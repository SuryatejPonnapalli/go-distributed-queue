package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct{
	Code int `json:"code"`
	Message string	`json:"message"`
}

func WriteError(c *gin.Context, code int, message string){
	c.JSON(code, ErrorResponse{
		Code: code,
		Message: message,
	})
}

var (
	RequestError = func(c *gin.Context, err error){
		WriteError(c, http.StatusBadRequest, err.Error())
	}

	InternalError = func(c *gin.Context) {
		WriteError(c, http.StatusInternalServerError, "An unexpected error occured.")
	}

	Unathorized = func(c *gin.Context){
		WriteError(c, http.StatusUnauthorized, "Unauthorized")
	}
)