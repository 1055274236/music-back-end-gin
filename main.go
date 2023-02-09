package main

import (
	"MusicBackEndOfGin/config"
	"MusicBackEndOfGin/middleware"
	"MusicBackEndOfGin/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create(config.GetConf().Gin.Logpath)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	middleware.InitMiddleware(engine) // 添加全局中间件
	router.InitRouter(engine)         // 设置路由
	engine.Run(config.GetConf().Gin.Port)
}
