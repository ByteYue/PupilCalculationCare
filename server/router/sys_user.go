package router

import (
	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", Login)
		UserRouter.POST("register", Register)
		UserRouter.GET("practice", PracticeGenerate)
		UserRouter.POST("DIY", PracticeDIY)
		UserRouter.POST("commit", Register)
	}
}
