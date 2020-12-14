package router

import (
	"my-app/expression"
	//"db"
	"github.com/gin-gonic/gin"
)

//TODO 给QIQI了
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

//TODO 给QIQI了
func Login(c *gin.Context) {
	//var R request.Login
	//检查数据库查看用户名和密码是否正确
}

func PracticeGenerate(c *gin.Context) {
	examlevel := expression.LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	msg := expression.StartExamine(examlevel)
	c.JSON(200, msg)
}

func PracticeDIY(c *gin.Context) {
	var ch expression.LevelChoose
	if c.ShouldBind(&ch) == nil {

	}
	msg := expression.StartExamine(ch)
	c.JSON(200, msg)
}

//TODO 给QIQI了
//TODO 和pzc确定Commit返回什么数据
func CommitHandler(c *gin.Context) {

}
