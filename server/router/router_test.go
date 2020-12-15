package router_test

import (
	"my-app/expression"
	"my-app/router"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinSentMessage(t *testing.T) {
	examlevel := expression.LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	msg := expression.StartExamine(examlevel, false)
	r := gin.Default()
	r.GET("/practice", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"work": "done",
		})
		c.JSON(200, msg)
	})
	r.Run(":8080")
}

func TestSendHtml(t *testing.T) {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.Run(":8080")
}

func TestCreate(t *testing.T) {
	r := gin.Default()
	r.POST("/create", router.PracticeDIY)
	r.Run(":8080")
}
