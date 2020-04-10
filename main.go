package main

import (
	"github.com/gin-gonic/gin"
	_ "speed/bootstrap"
	"speed/router"
)

func main() {

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/views/*")
	router.Router(engine) //初始化路由



	_ = engine.Run(":8086")

}
