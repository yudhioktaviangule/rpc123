package app

import (
	"net/http"
	"tele-dialog-hooks/app/routers"

	"github.com/gin-gonic/gin"
)

func RouterSetting() *gin.Engine {
	ENGINE := gin.Default()
	hook := ENGINE.Group("webhooks")
	product := ENGINE.Group("products")
	ENGINE.GET("/", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, gin.H{
			"app_name":      "ARAFAT CELL",
			"lang":          "GO-Lang",
			"email_support": "yudhioktaviangule@gmail.com",
		})
	})
	routers.ProductRouters(product)
	routers.HooksRouter(hook)
	return ENGINE
}
