package routers

import (
	"tele-dialog-hooks/controllers"

	"github.com/gin-gonic/gin"
)

func HooksRouter(rg *gin.RouterGroup) {
	rg.POST("/", controllers.TeleWebhooks)
}
