package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"Message": "Something went wrong!",
		"Error":   err.Error(),
	}
}

func SuccessResponse(res any) gin.H {
	return gin.H{
		"Message": "Successfull!",
		"Data":    res,
	}
}
