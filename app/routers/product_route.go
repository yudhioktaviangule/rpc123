package routers

import (
	"tele-dialog-hooks/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouters(rg *gin.RouterGroup) {
	rg.GET("/", controllers.ProductHttpGet)
}
