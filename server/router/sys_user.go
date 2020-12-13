package router

import (
	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", Register)
	}
}
