package router

import(
	"github.com/gin-gonic/gin"
)

func InitGetRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", Login)
		UserRouter.POST("register", Register)
	}
}