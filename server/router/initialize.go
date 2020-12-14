package router

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Routers = gin.Default()

	Routers.GET("/", func(c *gin.Context) {
		//将首页HTML传给CLIENT
		//TODO 等前端把资源文件发过来
	})
	v1 := Routers.Group("")
	InitUserRouter(v1)
	{
		InitGetRouter(v1)
		InitUserRouter(v1)
		//TODO出题,DIY,出分,查看信息
	}
	return Routers
}
