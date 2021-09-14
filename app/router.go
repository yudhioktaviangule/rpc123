package app

import (
	"tele-dialog-hooks/app/routers"

	"github.com/gin-gonic/gin"
)

func RouterSetting() *gin.Engine {
	ENGINE := gin.Default()
	hook := ENGINE.Group("webhooks")
	product := ENGINE.Group("products")
	routers.ProductRouters(product)
	routers.HooksRouter(hook)
	return ENGINE
}
