package controllers

import (
	"net/http"
	"tele-dialog-hooks/helpers"
	"tele-dialog-hooks/models"

	"github.com/gin-gonic/gin"
)

func ProductHttpGet(ctx *gin.Context) {
	list := models.ProductList()
	ctx.SecureJSON(http.StatusOK, helpers.ResponseSuccessWithData("OK", list))

}
