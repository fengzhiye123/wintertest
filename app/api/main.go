package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"winter_test/app/api/global"
	"winter_test/app/api/internal/initialize"
	"winter_test/app/api/router"
)

func main() {
	initialize.SetupViper()
	initialize.SetupLogger()
	initialize.SetupDataBase()
	config := global.Config.ServerConfig
	gin.SetMode(config.Mode)
	global.Logger.Info("init server success", zap.String("port", config.Port+":"+config.Port))
	err := router.InitRouter(config.Port)
	if err != nil {
		global.Logger.Fatal("server start up failed," + err.Error())
	}
}
