package main

import (
	"github.com/gin-gonic/gin"
)

type Exp struct {
	Times	string	`json:"Nums"`
	//Expree	expression.Ex	`json:"Exp"`
}

func main() {
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"ping",
		})
	})
	r.Run()
	//expression.GenerateWholeJsonExpression(expression.StartExamine(2,10,10))
	//ch:=expression.LevelChoose{4,10,10}
	//expression.StartExamine(ch)
	//expression.StartExamine(4,10,10)
}
