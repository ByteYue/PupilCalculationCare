package router

import (
	//"my-app/model/request"
	//"db"
	"github.com/gin-gonic/gin"
)

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	//var R request.Register
	//检查数据库查看 用户名是否被使用
	
}

func Login(c *gin.Context) {
	//var R request.Login
	//检查数据库查看用户名和密码是否正确
}
