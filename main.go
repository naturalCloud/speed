package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	app "speed/bootstrap"
	"speed/router"
)

func main() {

	var ginMode string

	if app.AppEnv == "prod" {
		ginMode = gin.ReleaseMode
		gin.DefaultWriter = ioutil.Discard
	} else {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/views/*")
	router.Router(engine) //初始化路由

	_ = engine.Run(":8086")

}
