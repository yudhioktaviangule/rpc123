package main

import (
	"fmt"
	"tele-dialog-hooks/app"
	"tele-dialog-hooks/config"
	"tele-dialog-hooks/models"
)

func Setup() {
	config.InitEnv()
	models.Migrate()
	engineConfig := app.RouterSetting()
	addr := fmt.Sprintf("%s:%s", "localhost", "8000")
	engineConfig.Run(addr)
}

func main() {
	Setup()
}
