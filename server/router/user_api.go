package router

import (
	"my-app/expression"
	//"db"
	"github.com/gin-gonic/gin"
)

//PracticeGenerate 生成Practice给client
func PracticeGenerate(c *gin.Context) {
	examlevel := expression.LevelChoose{ //set to default
		Level: 3,
		Max:   20,
		Nums:  20,
	}
	msg := expression.StartExamine(examlevel, false)
	c.JSON(200, msg)
}

//PracticeDIY 生成DIY给client
func PracticeDIY(c *gin.Context) {
	var ch expression.LevelChoose
	if c.ShouldBindJSON(&ch) != nil {

	}
	msg := expression.StartExamine(ch, true)
	c.JSON(200, msg)
}

//CommitHandler wait
//TODO 给QIQI了
//TODO 和pzc确定Commit返回什么数据
