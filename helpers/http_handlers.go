package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(message string) gin.H {
	fin := gin.H{
		"code":    200,
		"message": message,
	}
	return fin
}

func ResponseSuccessWithData(message string, data interface{}) gin.H {
	fin := gin.H{

		"queryResult": data,
		"pesan":       message,
	}
	return fin
}

func ResponseFail(message string) gin.H {
	fin := gin.H{
		"code":    http.StatusNotAcceptable,
		"message": message,
	}
	return fin
}

func ResponseNotAuthorize() gin.H {
	fin := gin.H{
		"code":    http.StatusForbidden,
		"message": "You dont have any access",
	}
	return fin
}
